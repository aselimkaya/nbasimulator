package api

import (
	"net/http"

	"github.com/aselimkaya/nbasimulator/src/simulator"
	"github.com/labstack/echo/v4"
)

type Score struct {
	Away      string
	AwayScore int
	Home      string
	HomeScore int
	Remaining int
}

func GetScores(c echo.Context) error {
	games := make([]Score, 0)
	for _, game := range simulator.Schedule {
		s := Score{
			Away:      game.Game.Away.Team.Abbreviation,
			AwayScore: game.Game.Away.TeamStats.Score,
			Home:      game.Game.Home.Team.Abbreviation,
			HomeScore: game.Game.Home.TeamStats.Score,
			Remaining: game.Duration,
		}

		games = append(games, s)
	}
	return c.Render(http.StatusOK, "results.html", map[string]interface{}{
		"games": games,
	})
}

func GetAssistLeader(c echo.Context) error {
	s := simulator.New()
	leader, err := s.PlayerGameInfoService.GetAssistLeader()
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", nil)
	}

	return c.Render(http.StatusOK, "leader.html", map[string]interface{}{
		"Name":    leader.Player.Name,
		"Team":    leader.Player.Team,
		"Assists": leader.PlayerStats.Assist,
	})
}
