package simulator

import (
	"context"
	"fmt"
	"os"

	"github.com/aselimkaya/nbasimulator/src/repository"
	"github.com/aselimkaya/nbasimulator/src/service"
)

type Simulator struct {
	TeamService   service.TeamService
	PlayerService service.PlayerService
}

func New() *Simulator {
	conn, err := repository.NewConnection(context.Background(), os.Getenv("MONGO_DATABASE_NAME"))
	if err != nil {
		fmt.Println(err)
	}
	return &Simulator{
		TeamService:   *service.NewTeamService(conn, os.Getenv("TEAM_COLLECTION")),
		PlayerService: *service.NewPlayerService(conn, os.Getenv("PLAYER_COLLECTION")),
	}
}

func (s *Simulator) Run() {
	s.initDB()

	game, err := s.ScheduleNewGame()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(game)
}
