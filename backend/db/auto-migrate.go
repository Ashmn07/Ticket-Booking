package db

import (
	"github.com/Ashmn07/ticket-booking/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
