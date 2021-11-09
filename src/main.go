package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"sync"

	"github.com/aselimkaya/nbasimulator/src/api"
	"github.com/aselimkaya/nbasimulator/src/simulator"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

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
		renderer := &TemplateRenderer{
			templates: template.Must(template.ParseGlob("views/*.html")),
		}
		e.Renderer = renderer

		e.GET("/results", api.GetScores)
		e.GET("/leaderboard", api.GetAssistLeader)
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
	}()

	wg.Wait()
}
