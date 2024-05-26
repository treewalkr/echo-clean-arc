package db

import (
	"echo-clean-arc/domain"
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once sync.Once
var db *gorm.DB

func New(dsn string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			err = fmt.Errorf("failed to connect database: %w", err)
		}

		// Migrate the schema
		db.AutoMigrate(&domain.User{})
	})
	return db, err
}
