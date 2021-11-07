package service

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameService struct {
	Conn     *repository.Connection
	CollName string
}

func NewGameService(conn *repository.Connection, collName string) *GameService {
	return &GameService{Conn: conn, CollName: collName}
}

func (s *GameService) FindByID(id string) (collection.Game, error) {
	game := collection.Game{}
	err := s.Conn.DB.Collection(s.CollName).FindOne(s.Conn.Ctx, collection.Game{GameID: id}).Decode(&game)

	if err != nil {
		return collection.Game{}, fmt.Errorf("game could not be found, given: %s, error: %s", id, err.Error())
	}

	return game, nil
}

func (s *GameService) Insert(p collection.Game) (*mongo.InsertOneResult, error) {
	res, err := s.Conn.DB.Collection(s.CollName).InsertOne(s.Conn.Ctx, p)
	if err != nil {
		return nil, fmt.Errorf("game could not be inserted, given: %v, error: %s", p, err.Error())
	}
	return res, nil
}

func (s *GameService) Delete(id string) error {
	res, err := s.Conn.DB.Collection(s.CollName).DeleteOne(s.Conn.Ctx, collection.Game{GameID: id})

	if err != nil {
		return fmt.Errorf("game could not be deleted, given: %s, error: %s", id, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %s", id)
	}

	return nil
}
