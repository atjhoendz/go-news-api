package main

import (
	"github.com/atjhoendz/go-news-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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
