package service

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamService struct {
	Conn     *repository.Connection
	CollName string
}

func NewTeamService(conn *repository.Connection, collName string) *TeamService {
	return &TeamService{Conn: conn, CollName: collName}
}

func (s *TeamService) FindByAbbreviation(abbr string) (collection.Team, error) {
	team := collection.Team{}
	err := s.Conn.DB.Collection(s.CollName).FindOne(s.Conn.Ctx, collection.Team{Abbreviation: abbr}).Decode(&team)

	if err != nil {
		return collection.Team{}, fmt.Errorf("player could not be found, given: %s, error: %s", abbr, err.Error())
	}

	return team, nil
}

func (s *TeamService) Insert(t collection.Team) (*mongo.InsertOneResult, error) {
	res, err := s.Conn.DB.Collection(s.CollName).InsertOne(s.Conn.Ctx, t)
	if err != nil {
		return nil, fmt.Errorf("team could not be inserted, given: %v, error: %s", t, err.Error())
	}
	return res, nil
}

func (s *TeamService) Delete(name string) error {
	res, err := s.Conn.DB.Collection(s.CollName).DeleteOne(s.Conn.Ctx, collection.Team{Name: name})

	if err != nil {
		return fmt.Errorf("team could not be deleted, given: %s, error: %s", name, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %s", name)
	}

	return nil
}
