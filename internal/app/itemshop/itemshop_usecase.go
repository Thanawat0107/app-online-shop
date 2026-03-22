package itemshop

import (
	"fmt"
	"log/slog"
	"math"

	"github.com/Thanawat0107/app-online-shop/internal/upload"
)

type itemShopUsecaseImpl struct {
	itemShopRepo ItemShopRepository
	imageBuilder upload.ImageBuilder
	logger       *slog.Logger
}

func NewItemShopUsecase(repo ItemShopRepository, imageBuilder upload.ImageBuilder, logger *slog.Logger) ItemShopUsecase {
	return &itemShopUsecaseImpl{
		itemShopRepo: repo,
		imageBuilder: imageBuilder,
		logger:       logger,
	}
}

func (u *itemShopUsecaseImpl) ItemList(req *RequestItemFilter) (*ResponseItemList, error) {
	list, total, err := u.itemShopRepo.Listing(req.Page, req.Limit, req.SearchText)
	fmt.Println(list, total, err)
	if err != nil {
		u.logger.Error("failed to get item list", "error", err)
		return nil, err
	}

	itemModels := make([]*ItemShopEntity, len(list))
	for i, item := range list {
		itemModels[i] = &ItemShopEntity{
			ItemId:      item.ItemId,
			Name:        item.Name,
			Description: item.Description,
			Picture:     u.imageBuilder.Build(item.Picture),
			Price:       item.Price,
		}
	}

	return &ResponseItemList{
		Items: itemModels,
		Pagination: &Pagination{
			Total:      int(total),
			TotalPages: int(math.Ceil(float64(total) / float64(req.Limit))),
		},
	}, nil
}
