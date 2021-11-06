package collection

type Team struct {
	ID      string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string   `json:"name,omitempty" bson:"name,omitempty"`
	Players []Player `json:"players,omitempty" bson:"players,omitempty"`
}
