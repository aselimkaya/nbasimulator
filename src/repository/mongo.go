package repository

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Ctx    context.Context
	Client *mongo.Client
	DB     *mongo.Database
}

func NewConnection(ctx context.Context, dbName string) (Connection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("SERVER_HOST"), os.Getenv("MONGO_PORT"))))
	if err != nil {
		return Connection{}, err
	}
	return Connection{Client: client, Ctx: ctx, DB: client.Database(dbName)}, nil
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
