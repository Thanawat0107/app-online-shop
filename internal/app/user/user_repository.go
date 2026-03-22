package user

import (
	"log/slog"

	"github.com/jinzhu/copier"

	"github.com/Thanawat0107/app-online-shop/config"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database"
	"github.com/Thanawat0107/app-online-shop/internal/infra/database/models"
)

type userRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewUserRepository(logger *slog.Logger, conf *config.Config) UserRepository {
	return &userRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *userRepositoryImpl) Create(user *UserEntity) (*UserEntity, error) {
	newUser := new(models.UserRecord)
	copier.Copy(newUser, user)

	userRecord := new(models.UserRecord)
	err := r.db.Connect().Create(newUser).Scan(userRecord).Error
	if err != nil {
		r.logger.Error("Failed to create user", "error", err)
		return nil, err
	}

	result := new(UserEntity)
	copier.Copy(result, userRecord)
	return result, nil
}

func (r *userRepositoryImpl) FindById(userId string) (*UserEntity, error) {
	userRecord := new(models.UserRecord)

	err := r.db.Connect().Where("UserId = ?", userId).First(userRecord).Error
	if err != nil {
		r.logger.Error("Failed to find user", "error", err)
		return nil, err
	}

	result := new(UserEntity)
	copier.Copy(result, userRecord)
	return result, nil
}
