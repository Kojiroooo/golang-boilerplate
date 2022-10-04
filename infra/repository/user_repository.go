package repository

import (
	"app/domain/dto"
	"app/domain/model"
	"app/domain/repository"
	"app/infra"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
	// user model.User
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryGorm{db: *db}
}

func (repo *UserRepositoryGorm) Create(authRequest dto.AuthRequest) model.User {

	user := userFromAuthRequest(authRequest)
	db := infra.Db.GetConnection()
	db.Create(&user)

	return user
}

func (repo *UserRepositoryGorm) FindById(userId int) model.User {

	user := model.User{}

	return user
}

func (repo *UserRepositoryGorm) GetAllPosts() []model.User {
	users := []model.User{}

	return users
}

func userFromAuthRequest(authRequest dto.AuthRequest) model.User {
	return model.User{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}
}
