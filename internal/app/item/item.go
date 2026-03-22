package item

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(app *echo.Echo, httpHandler ItemHttpHandler) {
	v1 := app.Group("/api/v1/items")

	v1.POST("/create", httpHandler.Create)
	v1.PUT("/edit", httpHandler.Edit)
	v1.DELETE("/delete/:item_id", httpHandler.Delete)
}
