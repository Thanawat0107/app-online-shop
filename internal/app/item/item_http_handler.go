package item

import (
	"errors"
	"strconv"

	"github.com/Thanawat0107/app-online-shop/internal/response"
	"github.com/labstack/echo/v5"
)

type itemHttpHandlerImpl struct {
	itemUsecase ItemUsecase
}

func NewItemHttpHandler(itemUsecase ItemUsecase) ItemHttpHandler {
	return &itemHttpHandlerImpl{
		itemUsecase: itemUsecase,
	}
}

func (h *itemHttpHandlerImpl) Create(pctx *echo.Context) error {
	req := new(RequestItemCreate)
	if err := pctx.Bind(req); err != nil {
		return response.BadRequest(pctx, err)
	}
	if err := pctx.Validate(req); err != nil {
		return response.BadRequest(pctx, err)
	}

	result, err := h.itemUsecase.CreateItem(req, "1")
	if err != nil {
		return response.BadRequest(pctx, err)
	}

	return response.Success(pctx, "success to create item", result)
}

func (h *itemHttpHandlerImpl) Edit(pctx *echo.Context) error {
	req := new(RequestItemEdit)
	if err := pctx.Bind(req); err != nil {
		return response.BadRequest(pctx, err)
	}
	if err := pctx.Validate(req); err != nil {
		return response.BadRequest(pctx, err)
	}

	result, err := h.itemUsecase.EditItem(req)
	if err != nil {
		return response.BadRequest(pctx, err)
	}

	return response.Success(pctx, "success to edit item", result)
}

func (h *itemHttpHandlerImpl) Delete(pctx *echo.Context) error {
	p_item_id := pctx.Param("item_id")
	if p_item_id == "" {
		return response.BadRequest(pctx, errors.New("Item id is required"))
	}
	itemId, err := strconv.Atoi(p_item_id)
	if err != nil {
		return response.BadRequest(pctx, err)
	}

	if err := h.itemUsecase.DeleteItem(itemId); err != nil {
		return response.BadRequest(pctx, err)
	}

	return response.Success(pctx, "success to delete item", nil)
}
