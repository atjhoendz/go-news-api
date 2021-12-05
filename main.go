package main

import (
	"github.com/atjhoendz/go-news-api/database"
	"github.com/atjhoendz/go-news-api/news/models"
	"github.com/atjhoendz/go-news-api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	db := database.GetInstance()
	err = db.AutoMigrate(&models.News{})

	if err != nil {
		log.Fatalf("Migration error %v", err)
	}

	api := echo.New()
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${method}] uri=${uri} status=${status}\n",
	}))
	api.Use(middleware.Recover())

	routes.DefineApiRoute(api)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		//if req.URL.Path[:4] == "/api" {
		api.ServeHTTP(res, req)
		//}
		return
	})
	server.Logger.Fatal(server.Start(":3000"))
}
