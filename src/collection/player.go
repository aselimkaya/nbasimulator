package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"username,omitempty" bson:"username,omitempty"`
	Team string             `json:"team,omitempty" bson:"team,omitempty"`
}
