package database

import (
	"pokeShowcase-api/configs"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getConnection(dsn string) (db *gorm.DB, err error) {
	if configs.IsApplicationInTestMode() {
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
