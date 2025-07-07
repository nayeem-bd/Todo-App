package migrations

import (
	"github.com/nayeem-bd/Todo-App/domain"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Todo{})
	if err != nil {
		logger.Fatal("Failed to migrate database:", err)
		return
	}
}
