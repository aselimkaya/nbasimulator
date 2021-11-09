package simulator

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"github.com/aselimkaya/nbasimulator/src/service"
)

var Schedule []collection.ScheduledGame

type Simulator struct {
	GameService           service.GameService
	TeamService           service.TeamService
	PlayerService         service.PlayerService
	PlayerGameInfoService service.PlayerGameInfoService
}

func init() {
	Schedule = make([]collection.ScheduledGame, 0)
}

func New() *Simulator {
	conn, err := repository.NewConnection(context.Background(), os.Getenv("MONGO_DATABASE_NAME"))
	if err != nil {
		fmt.Println(err)
	}
	return &Simulator{
		GameService:           *service.NewGameService(conn, os.Getenv("GAME_COLLECTION")),
		TeamService:           *service.NewTeamService(conn, os.Getenv("TEAM_COLLECTION")),
		PlayerService:         *service.NewPlayerService(conn, os.Getenv("PLAYER_COLLECTION")),
		PlayerGameInfoService: *service.NewPlayerGameInfoService(conn, os.Getenv("PLAYER_GAME_INFO_COLLECTION")),
	}
}

func (s *Simulator) Run() {
	schedule, err := s.setWeeklySchedule()
	if err != nil {
		fmt.Println(err)
		return
	}

	Schedule = schedule

	var wg sync.WaitGroup
	wg.Add(len(schedule))

	for i := 0; i < len(schedule); i++ {
		go func(i int) {
			defer wg.Done()
			s.runSingleGame(&schedule[i])
		}(i)
	}
	wg.Wait()
}

func (s *Simulator) runSingleGame(game *collection.ScheduledGame) {
	attacker, defender := setAttackOrder(&game.Game.Away, &game.Game.Home)
	remainingAttacks := 2

	wg := sync.WaitGroup{}

	for game.Duration > 0 {
		//Attack starting
		for remainingAttacks > 0 {
			wg.Add(2)
			go func() { //sleepy branch that ensures passing of time
				defer wg.Done()
				time.Sleep(time.Second)
			}()

			go func() { // worker branch
				defer wg.Done()

				if isSuccessfulAttack() {
					//Player is going to shoot
					point, success := shoot()
					scorer, assister := getPlayersWhoScoreAndAssist(len(attacker.Players))
					if success {
						//Successful shoot from player
						if point == 2 {
							attacker.Players[scorer].PlayerStats.TwoPointMade++
							attacker.Players[scorer].PlayerStats.TwoPointAttempt++
						} else if point == 3 {
							attacker.Players[scorer].PlayerStats.ThreePointMade++
							attacker.Players[scorer].PlayerStats.ThreePointAttempt++
						}

						attacker.TeamStats.Score += point
						attacker.Players[assister].PlayerStats.Assist++

						remainingAttacks = 0
					} else {
						//Player missed the shoot
						if point == 2 {
							attacker.Players[scorer].PlayerStats.TwoPointAttempt++
						} else if point == 3 {
							attacker.Players[scorer].PlayerStats.ThreePointAttempt++
						}

						remainingAttacks--
					}
				} else {
					//Attack attempt failed
					remainingAttacks--
				}

				game.Duration--
			}()

			wg.Wait()
		}

		attacker, defender = defender, attacker
		remainingAttacks = 2
	}

	_, err := s.GameService.Insert(game.Game)
	if err != nil {
		fmt.Printf("Game %s could not be inserted into db. Error: %s", game.Game.GameID, err.Error())
	}

	for _, a := range game.Game.Away.Players {
		_, err := s.PlayerGameInfoService.Insert(a)
		if err != nil {
			fmt.Printf("Game info for player %s and game %s could not be inserted into db. Error: %s", a.Player.Name, a.GameID, err.Error())
		}
	}

	for _, h := range game.Game.Home.Players {
		_, err := s.PlayerGameInfoService.Insert(h)
		if err != nil {
			fmt.Printf("Game info for player %s and game %s could not be inserted into db. Error: %s", h.Player.Name, h.GameID, err.Error())
		}
	}
}
