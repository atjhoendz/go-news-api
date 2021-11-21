package models

import "github.com/atjhoendz/go-news-api/common/models"

type News struct {
	models.Base
	Title    string
	Content  string
	AuthorID int
}
