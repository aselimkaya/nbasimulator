package simulator

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
)

func (s *Simulator) scheduleNewGame(away, home collection.Team) (collection.ScheduledGame, error) {
	gameID := fmt.Sprintf("%svs%s", away.Abbreviation, home.Abbreviation)

	awayInfo, homeInfo := prepareTeamGameInfo(away, gameID), prepareTeamGameInfo(home, gameID)

	return collection.ScheduledGame{
		Game: collection.Game{
			GameID: gameID,
			Away:   awayInfo,
			Home:   homeInfo,
		},
		Duration: 240,
	}, nil

}

func (s *Simulator) setWeeklySchedule() ([]collection.ScheduledGame, error) {
	teams, err := s.TeamService.GetAllTeams()
	if err != nil {
		return nil, err
	}

	weeklySchedule := make([]collection.ScheduledGame, 0)

	for len(teams) > 0 {
		awayIndex := generateRandomNumber(len(teams))
		away := teams[awayIndex]
		teams = remove(teams, awayIndex) // remove selected team

		homeIndex := generateRandomNumber(len(teams))
		home := teams[homeIndex]
		teams = remove(teams, homeIndex)

		game, err := s.scheduleNewGame(away, home)
		if err != nil {
			return nil, fmt.Errorf("%s vs %s could not be scheduled due to error: %s", away.Abbreviation, home.Abbreviation, err.Error())
		}
		weeklySchedule = append(weeklySchedule, game)
	}

	return weeklySchedule, nil
}

func prepareTeamGameInfo(team collection.Team, gameID string) collection.TeamGameInfo {
	players := make([]collection.PlayerGameInfo, 0)

	for _, p := range team.Players {
		players = append(players, preparePlayerGameInfo(p, gameID))
	}

	return collection.TeamGameInfo{
		GameID:  gameID,
		Team:    team,
		Players: players,
		TeamStats: collection.TeamStats{
			GameID: gameID,
		},
	}
}

func preparePlayerGameInfo(player collection.Player, gameID string) collection.PlayerGameInfo {
	return collection.PlayerGameInfo{
		GameID: gameID,
		Player: player,
	}
}

func remove(slice []collection.Team, s int) []collection.Team {
	return append(slice[:s], slice[s+1:]...)
}
