package service

import (
	"context"
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamService struct {
	Conn     *repository.Connection
	CollName string
}

func NewTeamService(conn *repository.Connection, collName string) *TeamService {
	return &TeamService{Conn: conn, CollName: collName}
}

func (s *TeamService) GetAllTeams() ([]collection.Team, error) {
	teams := make([]collection.Team, 0)
	cursor, err := s.Conn.DB.Collection(s.CollName).Find(s.Conn.Ctx, bson.D{})

	if err != nil {
		return []collection.Team{}, fmt.Errorf("all teams could not be retrieved, error: %s", err.Error())
	}

	for cursor.Next(context.TODO()) {
		t := collection.Team{}
		err := cursor.Decode(&t)
		if err != nil {
			fmt.Println(err)
		}

		teams = append(teams, t)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println(err)
	}

	cursor.Close(s.Conn.Ctx)

	return teams, nil
}

func (s *TeamService) FindByAbbreviation(abbr string) (collection.Team, error) {
	team := collection.Team{}
	err := s.Conn.DB.Collection(s.CollName).FindOne(s.Conn.Ctx, collection.Team{Abbreviation: abbr}).Decode(&team)

	if err != nil {
		return collection.Team{}, fmt.Errorf("team could not be found, given: %s, error: %s", abbr, err.Error())
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
