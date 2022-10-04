package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func ConnectDB() (*gorm.DB, error) {
	var db *gorm.DB

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
