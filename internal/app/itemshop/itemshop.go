package itemshop

import (
	"net/http"

	"github.com/Thanawat0107/app-online-shop/config"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database/models"
	"github.com/Thanawat0107/app-online-shop/internal/response"

	"github.com/labstack/echo/v5"
)

type ItemShopHandler struct {
	conf *config.Config
}

func NewItemShopHandler(conf *config.Config) *ItemShopHandler {
	return &ItemShopHandler{
		conf: conf,
	}
}

func (h *ItemShopHandler) GetItemShop(pctx *echo.Context) error {
	var items []models.ItemRecord

	db := h.conf.GetDb("mssql")
	if result := db.Connect().Find(&items); result.Error != nil {
		return pctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: "Failed to fetch items",
			Error:   result.Error.Error(),
		})
	}

	return pctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "Items fetched successfully",
		Data:    items,
	})
}
