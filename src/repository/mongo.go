package repository

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aselimkaya/nbasimulator/src/collection"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Ctx    context.Context
	Client *mongo.Client
	DB     *mongo.Database
}

func NewConnection(ctx context.Context, dbName string) (Connection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s", os.Getenv("MONGO_SCHEMA"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))))
	if err != nil {
		return Connection{}, err
	}

	conn := Connection{Client: client, Ctx: ctx, DB: client.Database(dbName)}
	conn.mongoInit(dbName)

	return conn, nil
}

func (c *Connection) NewCollection(collName string) (*mongo.Collection, error) {
	err := c.DB.CreateCollection(c.Ctx, collName)
	if err != nil {
		if strings.Contains(err.Error(), "Collection already exists") {
			return c.GetCollection(collName), nil
		} else {
			return nil, err
		}
	}
	return c.GetCollection(collName), nil
}

func (c *Connection) GetCollection(collName string) *mongo.Collection {
	return c.DB.Collection(collName)
}

func (c *Connection) mongoInit(dbName string) {
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

	players := []collection.Player{p1, p2, p3, p4, p5}

	for _, p := range players {
		_, err := c.DB.Collection("player").InsertOne(c.Ctx, p)
		if err != nil {
			fmt.Println(err)
		}
	}

	t1 := collection.Team{
		ID:      "BOS",
		Name:    "Boston Celtics",
		Players: players,
	}

	_, err := c.DB.Collection("team").InsertOne(c.Ctx, t1)
	if err != nil {
		fmt.Println(err)
	}
}
