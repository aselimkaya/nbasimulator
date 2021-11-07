package collection

type Game struct {
	GameID string       `json:"game_id,omitempty" bson:"game_id,omitempty"`
	Away   TeamGameInfo `json:"away,omitempty" bson:"away,omitempty"`
	Home   TeamGameInfo `json:"home,omitempty" bson:"home,omitempty"`
}

type ScheduledGame struct {
	Game     Game `json:"game,omitempty" bson:"game,omitempty"`
	Duration int  `json:"duration,omitempty" bson:"duration,omitempty"`
}
