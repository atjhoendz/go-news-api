package news

import (
	"github.com/atjhoendz/go-news-api/common"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type (
	Controller struct {
		NewsService Service
	}

	AddNewsRequestBody struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
		AuthorID int    `json:"authorID"`
	}
)

func (c Controller) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/news",
			Handler: c.AddNewsHandler,
		},
		{
			Method:  echo.GET,
			Path:    "/news",
			Handler: c.GetAllNewsHandler,
		},
		{
			Method:  echo.GET,
			Path:    "/news/:id",
			Handler: c.GetOneNewsHandler,
		},
		{
			Method:  echo.DELETE,
			Path:    "/news/:id",
			Handler: c.RemoveNewsHandler,
		},
	}
}

func (c Controller) AddNewsHandler(ctx echo.Context) error {
	body := new(AddNewsRequestBody)

	if err := ctx.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result := c.NewsService.AddNews(*body)

	return ctx.JSON(http.StatusCreated, result)
}

func (c Controller) GetAllNewsHandler(ctx echo.Context) error {
	result := c.NewsService.GetAll()

	return ctx.JSON(http.StatusOK, result)
}

func (c Controller) GetOneNewsHandler(ctx echo.Context) error {
	paramId := ctx.Param("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := c.NewsService.GetOne(id)

	if result == nil {
		return ctx.JSON(http.StatusNotFound, result)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c Controller) RemoveNewsHandler(ctx echo.Context) error {
	paramId := ctx.Param("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	result := c.NewsService.Remove(id)

	if result == nil {
		return ctx.JSON(http.StatusNotFound, result)
	}

	return ctx.JSON(http.StatusOK, result)
}
