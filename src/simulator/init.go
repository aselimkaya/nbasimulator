package simulator

import (
	"encoding/json"
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/go-resty/resty/v2"
)

type PlayerResp struct {
	Data []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Team      struct {
			Abbreviation string `json:"abbreviation"`
			FullName     string `json:"full_name"`
		} `json:"team"`
	} `json:"data"`

	Meta struct {
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	} `json:"meta"`
}

func (s *Simulator) InitFromAPI() {
	client := resty.New()
	page := 1
	fetchMore := true

	playerMap := make(map[string][]collection.Player)
	teamMap := make(map[string]string)

	for fetchMore {
		resp, err := client.R().
			SetQueryParams(map[string]string{
				"page":     fmt.Sprint(page),
				"per_page": "100",
			}).
			Get("https://www.balldontlie.io/api/v1/players")

		if err != nil {
			fmt.Println("player info could not be retrieved from api")
			return
		} else if resp.IsError() {
			fmt.Println("player request failed!", resp.Error())
			return
		}

		var r PlayerResp
		json.Unmarshal(resp.Body(), &r)

		for _, p := range r.Data {
			player := collection.Player{
				Name: fmt.Sprintf("%s %s", p.FirstName, p.LastName),
				Team: p.Team.Abbreviation,
			}

			if val, ok := playerMap[p.Team.Abbreviation]; ok {
				if len(val) <= 15 {
					playerMap[p.Team.Abbreviation] = append(val, player)
				}
			} else {
				playerMap[p.Team.Abbreviation] = []collection.Player{player}
			}

			if _, ok := teamMap[p.Team.Abbreviation]; !ok {
				teamMap[p.Team.Abbreviation] = p.Team.FullName
			}

		}

		if r.Meta.CurrentPage == r.Meta.TotalPages {
			fetchMore = false
		} else {
			page++
		}
	}

	for k, v := range teamMap {
		team := collection.Team{
			Abbreviation: k,
			Name:         v,
			Players:      playerMap[k],
		}

		_, err := s.TeamService.Insert(team)
		if err != nil {
			fmt.Printf("Team %s could not be inserted. Error: %s\n", team.Name, err.Error())
		}

		for _, p := range playerMap[k] {
			_, err := s.PlayerService.Insert(p)
			if err != nil {
				fmt.Printf("Player %s could not be inserted. Error: %s\n", p.Name, err.Error())
			}
		}
	}
}
