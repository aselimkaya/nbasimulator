package simulator

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
)

func (s *Simulator) ScheduleNewGame() (collection.ScheduledGame, error) {
	away, err := s.TeamService.FindByAbbreviation("DAL")
	if err != nil {
		return collection.ScheduledGame{}, fmt.Errorf("internal server error! away team could not retrieved from database: %s", err.Error())
	}

	home, err := s.TeamService.FindByAbbreviation("BOS")
	if err != nil {
		return collection.ScheduledGame{}, fmt.Errorf("internal server error! home team could not retrieved from database: %s", err.Error())
	}

	gameID := fmt.Sprintf("%svs%s", away.Abbreviation, home.Abbreviation)

	awayInfo, homeInfo := prepareTeamGameInfo(away, gameID), prepareTeamGameInfo(home, gameID)

	return collection.ScheduledGame{
		Game: collection.Game{
			GameID: gameID,
			Away:   awayInfo,
			Home:   homeInfo,
		},
		Duration:      240,
		AttackingTeam: &homeInfo,
	}, nil

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
		PlayerStats: collection.PlayerStats{
			GameID: gameID,
			Name:   player.Name,
		},
	}
}
