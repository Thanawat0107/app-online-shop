package item

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/Thanawat0107/app-online-shop/config"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database/models"
	"github.com/jinzhu/copier"
)

type itemRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewItemRepository(logger *slog.Logger, conf *config.Config) ItemRepository {
	return &itemRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *itemRepositoryImpl) FindById(itemId int) (*ItemEntity, error) {
	return nil, nil
}

func (r *itemRepositoryImpl) FindExists(itemId int) bool {
	var count int64
	err := r.db.Connect().Model(&models.ItemRecord{}).Where("ItemId = ?", itemId).Count(&count).Error
	if err != nil {
		r.logger.Error("Failed to fetch item", "error", err)
		return false
	}
	return count > 0
}

func (r *itemRepositoryImpl) Create(item *ItemEntity) (*ItemEntity, error) {
	newItem := new(models.ItemRecord)
	copier.Copy(newItem, item)

	itemRecord := new(models.ItemRecord)
	if err := r.db.Connect().Create(newItem).Scan(itemRecord).Error; err != nil {
		r.logger.Error("failed to create item", "error", err)
		return nil, err
	}

	result := new(ItemEntity)
	copier.Copy(result, itemRecord)
	return result, nil
}

func (r *itemRepositoryImpl) Edit(item *ItemEntity) (*ItemEntity, error) {
	updateItem := new(models.ItemRecord)
	copier.Copy(updateItem, item)

	if err := r.db.Connect().Where("ItemId = ?", updateItem.ItemId).Updates(updateItem).Error; err != nil {
		r.logger.Error("Failed to update item", "error", err)
		return nil, err
	}

	itemRecord := new(models.ItemRecord)
	if err := r.db.Connect().Where("ItemId = ?", item.ItemId).First(itemRecord).Error; err != nil {
		r.logger.Error("Failed to fetching item after update", "error", err)
		return nil, err
	}

	result := new(ItemEntity)
	copier.Copy(result, itemRecord)
	return result, nil
}

func (r *itemRepositoryImpl) Archive(itemId int) error {
	r.db.Connect().Model(&models.ItemRecord{}).Where("ItemId = ")
	err := r.db.Connect().Model(&models.ItemRecord{}).Where("ItemId = ?", itemId).Update("ActiveStatus", "UNAVAILABLE").Error
	if err != nil {
		r.logger.Error("Failed to archive item", "error", err)
		return err
	}
	return nil
}

func (r *itemRepositoryImpl) Begin() *gorm.DB {
	return nil
}

func (r *itemRepositoryImpl) Commit(tx *gorm.DB) error {
	return nil
}

func (r *itemRepositoryImpl) Rollback(tx *gorm.DB) error {
	return nil
}
