package news

import (
	"github.com/atjhoendz/go-news-api/database"
	"github.com/atjhoendz/go-news-api/news/models"
	"log"
)

type Service struct{}

func (s Service) AddNews(body AddNewsRequestBody) models.News {
	db := database.GetInstance()

	news := models.News{
		Title:    body.Title,
		Content:  body.Content,
		AuthorID: body.AuthorID,
	}

	db.Create(&news)

	return news
}

func (s Service) GetAll() []models.News {
	db := database.GetInstance()

	var news []models.News
	db.Find(&news)

	return news
}

func (s Service) GetOne(id int) *models.News {
	db := database.GetInstance()

	var news models.News
	result := db.First(&news, id)

	if result.RowsAffected < 1 {
		return nil
	}

	return &news
}

func (s Service) Remove(id int) *int64 {
	db := database.GetInstance()

	result := db.Delete(&models.News{}, id)

	if result.Error != nil {
		log.Printf("Delete error %v", result.Error)
		return nil
	}

	if result.RowsAffected < 1 {
		log.Printf("News with id %d doesn't exist\n", id)
		return nil
	}

	return &result.RowsAffected
}
