package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":80"))
}

func initRouting(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/api/go/hello", test)
}

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"go": "hello"})
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
