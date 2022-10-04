package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title  string
	Artist string
	Url    string
}