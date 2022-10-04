package controller

import (
	"app/domain/dto"
	"app/domain/model"
	"app/domain/service"

	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return AuthController{service: service}
}

func (controller *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var post model.Post
	err = json.Unmarshal(body[:length], &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var authRequest dto.AuthRequest
	err = json.Unmarshal(body[:length], &authRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	auth := controller.service.Signup(authRequest)

	json.NewEncoder(w).Encode(auth)
}

func (controller *AuthController) Signin(w http.ResponseWriter, r *http.Request) {
	controller.service.Signin()
}

func (controller *AuthController) Signout(w http.ResponseWriter, r *http.Request) {
	controller.service.Signout()
}
