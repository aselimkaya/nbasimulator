package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aselimkaya/nbasimulator/src/simulator"
	"github.com/labstack/echo/v4"
)

func main() {
	s := simulator.New()
	s.Run()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
