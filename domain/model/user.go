package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	DisplayName string
	Email       string
	Password    string
	Auths       []Auth
}

// user := model.User{DisplayName: "test", Email: "eeee@example.com", PassWord: "testtest", Token: "iajwfoiaje"}
// fmt.Println(user.DisplayName)
