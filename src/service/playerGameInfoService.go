package service

import (
	"fmt"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"github.com/aselimkaya/nbasimulator/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerGameInfoService struct {
	Conn     *repository.Connection
	CollName string
}

func NewPlayerGameInfoService(conn *repository.Connection, collName string) *PlayerGameInfoService {
	return &PlayerGameInfoService{Conn: conn, CollName: collName}
}

func (s *PlayerGameInfoService) Find(gameID, name string) (collection.PlayerGameInfo, error) {
	playerGameInfo := collection.PlayerGameInfo{}
	err := s.Conn.DB.Collection(s.CollName).FindOne(s.Conn.Ctx, bson.M{"game_id": gameID, "player.name": name}).Decode(&playerGameInfo)

	if err != nil {
		return collection.PlayerGameInfo{}, fmt.Errorf("player game info could not be found, given: %s, error: %s", name, err.Error())
	}

	return playerGameInfo, nil
}

func (s *PlayerGameInfoService) Insert(p collection.PlayerGameInfo) (*mongo.InsertOneResult, error) {
	res, err := s.Conn.DB.Collection(s.CollName).InsertOne(s.Conn.Ctx, p)
	if err != nil {
		return nil, fmt.Errorf("player game info could not be inserted, given: %v, error: %s", p, err.Error())
	}
	return res, nil
}

func (s *PlayerGameInfoService) Delete(gameID, name string) error {
	res, err := s.Conn.DB.Collection(s.CollName).DeleteOne(s.Conn.Ctx, bson.M{"game_id": gameID, "player.name": name})

	if err != nil {
		return fmt.Errorf("player game info could not be deleted, given: %s, error: %s", name, err.Error())
	} else if res.DeletedCount == 0 {
		return fmt.Errorf("no match found to delete, given: %s", name)
	}

	return nil
}
