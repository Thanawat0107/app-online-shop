package itemshop

import "github.com/labstack/echo/v5"

type ItemShopHttpHandler interface {
	GetAll(c *echo.Context) error
}

type ItemShopUsecase interface {
	ItemList(filter *RequestItemFilter) (*ResponseItemList, error)
}

type ItemShopRepository interface {
	Listing(page int, limit int, searchText string) ([]*ItemShopEntity, int64, error)
}
