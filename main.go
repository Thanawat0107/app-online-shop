package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/uwuluck23uwu/app-online-shop/config"
	"github/uwuluck23uwu/app-online-shop/internal/app/itemshop"
	"github/uwuluck23uwu/app-online-shop/internal/middleware"

	"github.com/labstack/echo/v5"
)

type Server struct {
	*echo.Echo
	conf *config.Config
}

func NewServer() *Server {
	conf := config.NewConfig()
	app := echo.New()

	return &Server{
		Echo: app,
		conf: conf,
	}
}

func (s *Server) InitRoutes() {
	// ItemShop routes
	itemShopHandler := itemshop.NewItemShopHandler(s.conf)
	s.GET("/api/v1/item-shop", itemShopHandler.GetItemShop)
}

func main() {
	server := NewServer()

	server.Use(middleware.RequestLogger(server.conf.Env))

	server.GET("/health", func(pctx *echo.Context) error {
		return pctx.JSON(http.StatusOK, map[string]string{
			"message": "server is running...",
		})
	})

	// Initialize routes
	server.InitRoutes()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         ":" + server.conf.Env.APP_PORT,
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, server.Echo); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
