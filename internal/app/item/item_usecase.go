package item

import (
	"errors"
	"log/slog"
)

type itemUsecaseImpl struct {
	logger   *slog.Logger
	itemRepo ItemRepository
}

func NewItemUsecase(logger *slog.Logger, itemRepo ItemRepository) ItemUsecase {
	return &itemUsecaseImpl{
		logger:   logger,
		itemRepo: itemRepo,
	}
}

func (u *itemUsecaseImpl) CreateItem(req *RequestItemCreate, adminId string) (*ItemEntity, error) {
	entity := req.ToEntity()
	entity.AdminId = "1"

	result, err := u.itemRepo.Create(entity)
	if err != nil {
		u.logger.Error("failed to create repo item", "error", err)
		return nil, err
	}

	return result, nil
}

func (u *itemUsecaseImpl) EditItem(req *RequestItemEdit) (*ItemEntity, error) {
	entity := req.ToEntity()

	result, err := u.itemRepo.Edit(entity)
	if err != nil {
		u.logger.Error("failed to update repo item", "error", err)
		return nil, err
	}

	return result, nil
}

func (u *itemUsecaseImpl) DeleteItem(itemId int) error {
	if !u.itemRepo.FindExists(itemId) {
		return errors.New("item does not exist")
	}

	err := u.itemRepo.Archive(itemId)
	if err != nil {
		u.logger.Error("Failed to send archive item", "error", err)
		return err
	}

	return nil
}
