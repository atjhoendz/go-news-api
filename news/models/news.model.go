package models

import "github.com/atjhoendz/go-news-api/common/models"

type News struct {
	models.Base
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"authorID"`
}
