package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"app/domain/model"
	"app/domain/repository"
)

type PostController interface {
	PostList(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	repo repository.PostRepository
}

func NewPostController(postRepository repository.PostRepository) PostController {
	return &postController{postRepository}
}

func (controller *postController) PostList(w http.ResponseWriter, r *http.Request) {
	posts := controller.repo.GetAllPosts()

	// error処理書きたい
	json.NewEncoder(w).Encode(posts)
}

func (controller *postController) CreatePost(w http.ResponseWriter, r *http.Request) {
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
	controller.repo.Create(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
