package main

import (
	"log"

	"github/uwuluck23uwu/app-online-shop/config"
	"github/uwuluck23uwu/app-online-shop/internal/infra/database"
	"github/uwuluck23uwu/app-online-shop/internal/infra/database/models"
)

func main() {
	env := config.NewEnv()

	db := database.NewSqlServerDatabase(env.DB_HOST, env.DB_USERNAME, env.DB_PASSWORD, env.DB_NAME)

	tx := db.Connect().Begin()

	tx.AutoMigrate(&models.UserRecord{})
	tx.AutoMigrate(&models.UserBalanceRecord{})
	tx.AutoMigrate(&models.ItemRecord{})
	tx.AutoMigrate(&models.InventoryRecord{})
	tx.AutoMigrate(&models.PurchaseHistoryRecord{})

	if err := tx.Commit().Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Fatal("Error: migration failed")
		}
		log.Fatal("Error: migration successfully")
	}
}
