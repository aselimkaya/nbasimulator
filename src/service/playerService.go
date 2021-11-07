package service

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerService struct {
	Conn     *repository.Connection
	CollName string
}

func NewPlayerService(conn *repository.Connection, collName string) *PlayerService {
	return &PlayerService{Conn: conn, CollName: collName}
}

func (s *PlayerService) FindByName(name string) (collection.Player, error) {
	player := collection.Player{}
	err := s.Conn.DB.Collection(s.CollName).FindOne(s.Conn.Ctx, collection.Player{Name: name}).Decode(&player)

	if err != nil {
		return collection.Player{}, fmt.Errorf("player could not be found, given: %s, error: %s", name, err.Error())
	}

	return player, nil
}

func (s *PlayerService) Insert(p collection.Player) (*mongo.InsertOneResult, error) {
	res, err := s.Conn.DB.Collection(s.CollName).InsertOne(s.Conn.Ctx, p)
	if err != nil {
		return nil, fmt.Errorf("player could not be inserted, given: %v, error: %s", p, err.Error())
	}
	return res, nil
}

func (s *PlayerService) Delete(name string) error {
	res, err := s.Conn.DB.Collection(s.CollName).DeleteOne(s.Conn.Ctx, collection.Player{Name: name})

	if err != nil {
		return fmt.Errorf("player could not be deleted, given: %s, error: %s", name, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %s", name)
	}

	return nil
}
