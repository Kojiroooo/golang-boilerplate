package repository

import (
	"app/domain/model"
	"app/infra"
)

type PostRepository interface {
	Create(*model.Post) model.Post
	FindById(int) model.Post
	GetAllPosts() []model.Post
}

type postRepositoryGorm struct {
	post model.Post
	// db を持たせてもいいかも
}

func NewPostRepository() PostRepository {
	return &postRepositoryGorm{}
}

func (repo *postRepositoryGorm) Create(post *model.Post) model.Post {
	db := infra.Db.GetConnection()
	// post := Post{Title: "title dayo", Artist: "artist dayo", Url: "https://google.com"}
	db.Create(&post)

	return *post
}

func (repo *postRepositoryGorm) FindById(id int) model.Post {
	db := infra.Db.GetConnection()
	db.First(&repo.post, id)

	return repo.post
}

func (repo *postRepositoryGorm) GetAllPosts() []model.Post {
	createPost() // 一時的にデータ作成
	db := infra.Db.GetConnection()
	posts := []model.Post{}
	// res := db.Find(&posts)
	db.Find(&posts)
	return posts
}

func createPost() {
	db := infra.Db.GetConnection()
	post := model.Post{Title: "title dayo", Artist: "artist dayo", Url: "https://google.com"}
	db.Create(&post)
}

func NewPostRepositoryMock() PostRepository {
	return &postRepositoryMock{}
}

type postRepositoryMock struct {
	posts []model.Post
	post  model.Post
}

func (repo *postRepositoryMock) Create(post *model.Post) model.Post {
	db := infra.Db.GetConnection()
	// post := Post{Title: "title dayo", Artist: "artist dayo", Url: "https://google.com"}
	db.Create(&post)

	return *post
}

func (repo *postRepositoryMock) FindById(id int) model.Post {
	db := infra.Db.GetConnection()
	db.First(&repo.post, id)

	return repo.post
}

func (repo *postRepositoryMock) GetAllPosts() []model.Post {
	repo.posts = []model.Post{
		{Title: "title dayo1", Artist: "artist dayo1", Url: "url dayo1"},
		{Title: "title dayo2", Artist: "artist dayo2", Url: "url dayo2"},
		{Title: "title dayo3", Artist: "artist dayo3", Url: "url dayo3"},
	}

	// db := Db.GetConnection()
	// posts := []Post{}
	// // res := db.Find(&posts)
	// db.Find(&posts)
	return repo.posts
}
