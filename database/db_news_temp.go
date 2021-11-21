package database

import "github.com/atjhoendz/go-news-api/news/models"

type (
	News map[int]*models.News
)

var (
	ListNews = make(News)
)

func (db News) AddNews(id int, news *models.News) *models.News {
	ListNews[id] = news
	return ListNews[id]
}

func (db News) GetAllNews() News {
	return ListNews
}

func (db News) GetOneNews(id int) *models.News {
	return ListNews[id]
}

func (db News) UpdateNews(id int, news *models.News) *models.News {
	ListNews[id] = news
	return ListNews[id]
}

func (db News) RemoveNews(id int) {
	delete(ListNews, id)
}
