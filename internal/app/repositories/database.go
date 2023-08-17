package repositories

import (
	"github.com/zatrano/zatrano/internal/app/models"
	"github.com/zatrano/zatrano/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabase() *gorm.DB {
	if db != nil {
		return db
	}

	cfg := config.LoadConfig()
	dsn := config.GetDBConnectionString(cfg)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return db
}
