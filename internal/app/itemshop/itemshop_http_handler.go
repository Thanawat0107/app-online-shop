package itemshop

import (
	"github.com/labstack/echo/v5"

	"github.com/Thanawat0107/app-online-shop/internal/response"
)

type itemShopHttpHandlerImpl struct {
	itemShopUsecase ItemShopUsecase
}

func NewItemShopHttpHandler(usecase ItemShopUsecase) ItemShopHttpHandler {
	return &itemShopHttpHandlerImpl{
		itemShopUsecase: usecase,
	}
}

func (h *itemShopHttpHandlerImpl) GetAll(pctx *echo.Context) error {
	req := &RequestItemFilter{
		Page:  1,
		Limit: 10,
	}
	if err := pctx.Bind(req); err != nil {
		pctx.Logger().Error("failed to bind request", "error", err)
		return response.BadRequest(pctx, err)
	}

	list, err := h.itemShopUsecase.ItemList(req)
	if err != nil {
		pctx.Logger().Error("failed to get item list", "error", err)
		return response.InternalError(pctx, err)
	}

	return response.Success(pctx, "Item list fetched successfully", list)
}
