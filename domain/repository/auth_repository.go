package repository

import (
	"app/domain/dto"
	"crypto/rand"
	"errors"
	"fmt"

	"app/domain/model"
	"app/infra"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Authenticate(authReqest dto.AuthRequest) (model.Auth, error)
	Authorize(token string) (model.Auth, error)
}

type authRepositoryGorm struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepositoryGorm{}
}

func (repo *authRepositoryGorm) Authenticate(authReqest dto.AuthRequest) (model.Auth, error) {
	// email と password を DBの値と突合する
	// okだったら tokenを発行し、tokenを返す
	// ngだったら errを返す
	db := infra.Db.GetConnection()
	var user model.User
	db.Where("email = ? AND password = ?", authReqest.Email, authReqest.Password).First(&user)

	token := createToken()

	auth := model.Auth{
		Token:  token,
		UserID: user.ID,
	}
	db.Create(&auth)

	return auth, nil
}

func (repo *authRepositoryGorm) Authorize(token string) (model.Auth, error) {
	// email と password を DBの値と突合する
	// okだったら tokenを発行し、tokenを返す
	// ngだったら errを返す
	db := infra.Db.GetConnection()
	var auth model.Auth
	// db.First(&auth, 2)
	// fmt.Println(auth.Token)
	fmt.Println(token)
	result := db.Where("token = ?", token).Last(&auth)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return auth, errors.New("Authorization error")
	}

	return auth, nil
}

func createToken() string {
	token, _ := makeRandomStr(30)
	return token
}

func makeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_=+[{]}!@#$%^*()"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

// https://hodalog.com/generate-bcrypt-hash-from-password-using-go/
func password_hash(nonhash string) string {
	return "xxxxxx"
}
