package database

import (
	models2 "github.com/atjhoendz/go-news-api/common/models"
	"github.com/atjhoendz/go-news-api/news/models"
	"testing"
	"time"
)

var (
	mockNews = models.News{
		Base: models2.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Title:    "Sample News Title",
		Content:  "Mock Content",
		AuthorID: 1,
	}
)

func TestNews_AddNews(t *testing.T) {
	listMockNews := News{}
	listMockNews.AddNews(1, &mockNews)

	n := len(ListNews)

	if n != 1 {
		t.Fatal("Add news failed")
	}
}

func TestNews_GetAllNews(t *testing.T) {
	listMockNews := News{}
	listMockNews.AddNews(mockNews.ID, &mockNews)
	newsData := listMockNews.GetAllNews()

	if len(newsData) != 1 {
		t.Fatal("Get all news must be return a data")
	}
}

func TestNews_GetOneNews(t *testing.T) {
	listMockNews := News{}
	listMockNews.AddNews(mockNews.ID, &mockNews)
	newsData := listMockNews.GetOneNews(1)

	if newsData == nil {
		t.Fatal("Get one news must return a data")
	}

	if newsData.ID != mockNews.ID {
		t.Fatalf("Get a news data must return ID = %d", mockNews.ID)
	}
}

func TestNews_GetOneNews_ifNotExist(t *testing.T) {
	listMockNews := News{}
	newsData := listMockNews.GetOneNews(mockNews.ID)

	if newsData != nil {
		t.Fatal("Must return nil")
	}
}

func TestNews_UpdateNews(t *testing.T) {
	listMockNews := News{}
	listMockNews.AddNews(mockNews.ID, &mockNews)

	existingNews := listMockNews.GetOneNews(mockNews.ID)
	existingNews.Title = "New Title"

	updatedNews := listMockNews.UpdateNews(mockNews.ID, existingNews)

	if updatedNews.Title != existingNews.Title {
		t.Fatalf("Updated news title must be %s", existingNews.Title)
	}

	if ListNews[mockNews.ID].Title != existingNews.Title {
		t.Fatalf("Updated news title must be %s", existingNews.Title)
	}
}

func TestNews_RemoveNews(t *testing.T) {
	listMockNews := News{}
	listMockNews.AddNews(mockNews.ID, &mockNews)

	listMockNews.RemoveNews(mockNews.ID)

	deletedNews := listMockNews.GetOneNews(mockNews.ID)

	if deletedNews != nil {
		t.Fatal("Deleted news must be nil")
	}
}
