package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"

	"github.com/Thanawat0107/app-online-shop/config"
	"github.com/Thanawat0107/app-online-shop/internal/app/auth"
	"github.com/Thanawat0107/app-online-shop/internal/app/item"
	"github.com/Thanawat0107/app-online-shop/internal/app/itemshop"
	"github.com/Thanawat0107/app-online-shop/internal/app/user"
	"github.com/Thanawat0107/app-online-shop/internal/echovalidator"
	"github.com/Thanawat0107/app-online-shop/internal/middleware"
	"github.com/Thanawat0107/app-online-shop/internal/upload"
)

type Server struct {
	*echo.Echo
	conf         *config.Config
	imageBuilder upload.ImageBuilder
}

func NewServer() *Server {
	conf := config.NewConfig()
	app := echo.New()

	app.Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app.Validator = echovalidator.NewValidator()

	app.Static("/uploads", "./uploads")

	imageBuilder := upload.NewImageBuilder(conf.Env.APP_HOST, "./uploads")

	return &Server{
		Echo:         app,
		conf:         conf,
		imageBuilder: imageBuilder,
	}
}

func main() {
	server := NewServer()

	server.Use(echoMiddleware.Recover())
	server.Use(middleware.RequestLogger(server.conf.Env))

	server.GET("/api/v1/health", server.healthCheck)
	server.POST("/api/v1/upload/image", server.uploadImage)
	server.GET("*", server.notFound)

	// Implement dependencies
	itemShopRepo := itemshop.NewItemShopRepository(server.conf, server.Logger)
	itemShopUsecase := itemshop.NewItemShopUsecase(itemShopRepo, server.imageBuilder, server.Logger)
	itemShopHandler := itemshop.NewItemShopHttpHandler(itemShopUsecase)
	itemRepo := item.NewItemRepository(server.Logger, server.conf)
	itemUsecase := item.NewItemUsecase(server.Logger, itemRepo)
	itemHttpHandler := item.NewItemHttpHandler(itemUsecase)
	userRepo := user.NewUserRepository(server.Logger, server.conf)
	authUsecase := auth.NewAuthGoogleUsecase(userRepo)
	authHandler := auth.NewAuthGoogleHandler(
		server.Logger,
		server.conf.Env,
		config.NewGoogleOAuth2Config(server.conf.Env),
		authUsecase,
	)

	itemshop.RegisterRoutes(server.Echo, itemShopHandler)
	item.RegisterRoutes(server.Echo, itemHttpHandler)
	auth.RegisterRoutes(server.Echo, authHandler)

	// Start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":" + server.conf.Env.APP_PORT,
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, server.Echo); err != nil {
		log.Fatal("Error starting server error", err)
	}
	log.Println("Server shutdown gracefully")
}

func (s *Server) healthCheck(pctx *echo.Context) error {
	return pctx.JSON(http.StatusOK, map[string]string{"message": "Server is running!"})
}

func (s *Server) notFound(pctx *echo.Context) error {
	return pctx.JSON(http.StatusNotFound, map[string]string{"error": "Not Found"})
}

func (s *Server) uploadImage(pctx *echo.Context) error {
	path, err := s.imageBuilder.SaveImage(pctx)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return pctx.JSON(http.StatusOK, map[string]string{"url": path, "fullPath": s.imageBuilder.Build(path)})
}
