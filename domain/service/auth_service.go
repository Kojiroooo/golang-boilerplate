package service

import (
	"app/domain/dto"
	"app/domain/model"
	"app/domain/repository"
	"fmt"
)

type AuthService interface {
	Signin()
	Signup(authReqest dto.AuthRequest) model.Auth
	Signout()
}

type authServiceEmail struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
}

func NewAuthService(ar repository.AuthRepository, ur repository.UserRepository) AuthService {
	return &authServiceEmail{authRepository: ar, userRepository: ur}
}

func (service *authServiceEmail) Signup(authReqest dto.AuthRequest) model.Auth {
	// user作成
	// user_idだけ受け取るでも良い。
	user := service.userRepository.Create(authReqest)

	auth, err := service.authRepository.Authenticate(authReqest)

	if err != nil {
		//  error
	}

	fmt.Println(user)

	return auth
}

func (service *authServiceEmail) Signin() {

}

func (service *authServiceEmail) Signout() {

}
