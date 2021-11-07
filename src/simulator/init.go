package simulator

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
)

func (s *Simulator) initDB() {
	p1 := collection.Player{
		Name: "Marcus Smart",
		Team: "BOS",
	}

	p2 := collection.Player{
		Name: "Jaylen Brown",
		Team: "BOS",
	}

	p3 := collection.Player{
		Name: "Jason Tatum",
		Team: "BOS",
	}

	p4 := collection.Player{
		Name: "Al Horford",
		Team: "BOS",
	}

	p5 := collection.Player{
		Name: "Robert Williams III",
		Team: "BOS",
	}

	p6 := collection.Player{
		Name: "Luka Doncic",
		Team: "DAL",
	}

	p7 := collection.Player{
		Name: "Tim Hardaway Jr.",
		Team: "DAL",
	}

	p8 := collection.Player{
		Name: "Dorain Finney-Smith",
		Team: "DAL",
	}

	p9 := collection.Player{
		Name: "Kristaps Porzingis",
		Team: "DAL",
	}

	p10 := collection.Player{
		Name: "Dwight Powell",
		Team: "DAL",
	}

	players := []collection.Player{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10}

	for _, p := range players {
		_, err := s.PlayerService.Insert(p)
		if err != nil {
			fmt.Println(err)
		}
	}

	t1 := collection.Team{
		Abbreviation: "BOS",
		Name:         "Boston Celtics",
		Players:      players[:5],
	}

	t2 := collection.Team{
		Abbreviation: "DAL",
		Name:         "Dallas Mavericks",
		Players:      players[5:],
	}

	teams := []collection.Team{t1, t2}

	for _, t := range teams {
		_, err := s.TeamService.Insert(t)
		if err != nil {
			fmt.Println(err)
		}
	}
}
