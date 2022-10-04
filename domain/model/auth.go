package model

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model

	Token     string
	ExpiredAt time.Time
	UserID    uint
}
