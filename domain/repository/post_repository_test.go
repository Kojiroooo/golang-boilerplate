package repository

import (
	"testing"

	"app/domain/model"
)

func TestGetAllPosts(t *testing.T) {
	posts := []model.Post{
		{Title: "title dayo1", Artist: "artist dayo1", Url: "url dayo1"},
		{Title: "title dayo2", Artist: "artist dayo2", Url: "url dayo2"},
		{Title: "title dayo3", Artist: "artist dayo3", Url: "url dayo3"},
	}
	// 件数
	var repo = NewPostRepository()
	posts = repo.GetAllPosts()

	got := len(posts)
	expected := 3

	if got != expected {
		t.Errorf("TestGetAllPosts = %d; want %d;", got, expected)
	}
}
