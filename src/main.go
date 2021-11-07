package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/aselimkaya/nbasimulator/src/simulator"
	"github.com/labstack/echo/v4"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		s := simulator.New()
		s.InitFromAPI()
		s.Run()
	}()

	go func() {
		defer wg.Done()

		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
	}()

	wg.Wait()
}
