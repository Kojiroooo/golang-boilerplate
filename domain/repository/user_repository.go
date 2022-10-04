package repository

import (
	"app/domain/dto"
	"app/domain/model"
	"app/infra"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(authRequest dto.AuthRequest) model.User
	FindById(int) model.User
	GetAllPosts() []model.User
}

type userRepositoryGorm struct {
	db   *gorm.DB
	user model.User
}

func NewUserRepository() UserRepository {
	return &userRepositoryGorm{}
}

func (repo *userRepositoryGorm) Create(authRequest dto.AuthRequest) model.User {

	user := userFromAuthRequest(authRequest)
	db := infra.Db.GetConnection()
	db.Create(&user)

	return user
}

func (repo *userRepositoryGorm) FindById(userId int) model.User {

	user := model.User{}

	return user
}

func (repo *userRepositoryGorm) GetAllPosts() []model.User {
	users := []model.User{}

	return users
}

func userFromAuthRequest(authRequest dto.AuthRequest) model.User {
	return model.User{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}
}
