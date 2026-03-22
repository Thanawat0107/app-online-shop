package itemshop

import (
	"log/slog"

	"github.com/jinzhu/copier"

	"github.com/Thanawat0107/app-online-shop/config"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database/models"
)

type itemshopRepositoryImpl struct {
	db     database.Database
	logger *slog.Logger
}

func NewItemShopRepository(conf *config.Config, logger *slog.Logger) ItemShopRepository {
	return &itemshopRepositoryImpl{
		db:     conf.GetDb("main"),
		logger: logger,
	}
}

func (r *itemshopRepositoryImpl) Listing(page int, limit int, searchText string) ([]*ItemShopEntity, int64, error) {
	itemRecords := make([]*models.ItemRecord, 0)

	query := r.db.Connect().Model(&models.ItemRecord{}).Where("ActiveStatus = ?", "AVAILABLE")
	if searchText != "" {
		searchText := "%" + searchText + "%"
		query = query.Where("Name LIKE ? OR Description LIKE ?", searchText, searchText)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		r.logger.Error("failed to get item total", "error", err)
		return nil, 0, err
	}

	query = query.Offset((page - 1) * limit).Limit(limit)
	if err := query.Find(&itemRecords).Error; err != nil {
		r.logger.Error("failed to get item listing", "error", err)
		return nil, 0, err
	}

	results := make([]*ItemShopEntity, 0)
	copier.Copy(&results, &itemRecords)
	return results, total, nil
}
