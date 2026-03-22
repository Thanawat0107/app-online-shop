package auth

import (
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/Thanawat0107/app-online-shop/config"
)

var (
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type authGoogleHandlerImpl struct {
	logger            *slog.Logger
	env               *config.Env
	oauth2Config      *config.Oauth2Config
	authGoogleUsecase AuthGoogleUsecase
}

func NewAuthGoogleHandler(logger *slog.Logger, env *config.Env, oauth2Config *config.Oauth2Config, authGoogleUsecase AuthGoogleUsecase) AuthGoogleHandler {
	return &authGoogleHandlerImpl{
		logger:            logger,
		env:               env,
		oauth2Config:      oauth2Config,
		authGoogleUsecase: authGoogleUsecase,
	}
}

func (h *authGoogleHandlerImpl) GoogleLogin(pctx *echo.Context) error {
	state := randomState()
	callBackUrl := pctx.QueryParam("callback_url")

	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.StateCookieName,
		Value:    state,
		Path:     "/",
		HttpOnly: true,
	})
	pctx.SetCookie(&http.Cookie{
		Name:     h.oauth2Config.AppClientRedirectKey,
		Value:    callBackUrl,
		Path:     "/",
		HttpOnly: true,
	})

	return pctx.Redirect(http.StatusFound, h.oauth2Config.GoogleOAuth2Config.AuthCodeURL(state))
}

func (h *authGoogleHandlerImpl) GoogleLoginCallBack(pctx *echo.Context) error {
	return nil
}

func (h *authGoogleHandlerImpl) Logout(pctx *echo.Context) error {
	return nil
}
