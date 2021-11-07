package collection

type Player struct {
	Name string `json:"username,omitempty" bson:"username,omitempty"`
	Team string `json:"team,omitempty" bson:"team,omitempty"`
}

type PlayerStats struct {
	GameID            string `json:"game_id,omitempty" bson:"game_id,omitempty"`
	Name              string `json:"username,omitempty" bson:"username,omitempty"`
	TwoPointMade      int    `json:"two_point_made,omitempty" bson:"two_point_made,omitempty"`
	TwoPointAttempt   int    `json:"two_point_attempt,omitempty" bson:"two_point_attempt,omitempty"`
	ThreePointMade    int    `json:"three_point_made,omitempty" bson:"three_point_made,omitempty"`
	ThreePointAttempt int    `json:"three_point_attempt,omitempty" bson:"three_point_attempt,omitempty"`
	Assist            int    `json:"assist,omitempty" bson:"assist,omitempty"`
}

type PlayerGameInfo struct {
	GameID      string      `json:"game_id,omitempty" bson:"game_id,omitempty"`
	Player      Player      `json:"player,omitempty" bson:"player,omitempty"`
	PlayerStats PlayerStats `json:"player_stats,omitempty" bson:"player_stats,omitempty"`
}
