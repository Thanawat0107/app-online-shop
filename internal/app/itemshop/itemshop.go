package itemshop

import "github.com/labstack/echo/v5"

func RegisterRoutes(app *echo.Echo, httpHandler ItemShopHttpHandler) {
	v1 := app.Group("/api/v1/item-shop")

	v1.GET("", httpHandler.GetAll)
}
