package models

import (
	constant "scrabble/app/models/const"
	grid "scrabble/app/models/grid"
	"time"
)

type GameOptions struct {
	PlayerName    string
	PlayerId      string
	Time          int
	Dictionary    string
	GameMode      constant.GameMode
	OpponentName  *string
	IsMultiplayer bool
}

type GameModeData struct {
	GameMode constant.GameMode
}

type GameInfo struct {
	PlayerTime int
	Dictionary string
}

type JoinMultiplayerOption struct {
	GameId     string
	PlayerInfo PlayerInfo
}

type Game struct {
	Id               string
	Creator          Player
	Opponent         *Player
	Dict             string
	Grid             []grid.Node
	LetterReserve    []string
	Time             int
	Timer            int
	Capacity         int
	SkipCounter      int
	HasEnded         bool
	Mode             string
	IsMultiplayer    bool
	PublicObjectives *[]Objective
	StartedTime      *time.Time
	TotalTime        int
}

type ReconnectionInfo struct {
	Id   string
	Game Game
}
type GameService struct {
	GameSessions []Game
}
