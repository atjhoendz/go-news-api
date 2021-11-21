package news

import (
	models2 "github.com/atjhoendz/go-news-api/common/models"
	"github.com/atjhoendz/go-news-api/database"
	"github.com/atjhoendz/go-news-api/news/models"
	"time"
)

type Service struct {
}

var tableNews database.News

func (s Service) AddNews(body AddNewsRequestBody) *models.News {
	news := models.News{
		Base: models2.Base{
			ID:        body.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:    body.Title,
		Content:  body.Content,
		AuthorID: body.AuthorID,
	}

	createdNews := tableNews.AddNews(news.ID, &news)

	return createdNews
}

func (s Service) GetAll() database.News {
	return tableNews.GetAllNews()
}

func (s Service) GetOne(id int) *models.News {
	return tableNews.GetOneNews(id)
}

func (s Service) Remove(id int) *models.News {
	news := s.GetOne(id)

	if news == nil {
		return news
	}

	tableNews.RemoveNews(id)
	return news
}
