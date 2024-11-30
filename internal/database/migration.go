package database

import (
	"loyalty_central/internal/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Admin{},
		&models.Customer{},
		&models.Menu{},
		&models.Purchase{},
		&models.PurchaseDetail{},
	)
}
