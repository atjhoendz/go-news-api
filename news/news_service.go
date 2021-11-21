package news

import (
	models2 "github.com/atjhoendz/go-news-api/common/models"
	"github.com/atjhoendz/go-news-api/database"
	"github.com/atjhoendz/go-news-api/news/models"
	"time"
)

type Service struct {
	TableNews database.News
}

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

	createdNews := s.TableNews.AddNews(news.ID, &news)

	return createdNews
}

func (s Service) GetAll() database.News {
	return s.TableNews.GetAllNews()
}

func (s Service) GetOne(id int) *models.News {
	return s.TableNews.GetOneNews(id)
}

func (s Service) Remove(id int) *models.News {
	news := s.GetOne(id)

	if news == nil {
		return news
	}

	s.TableNews.RemoveNews(id)
	return news
}
