package models

import (
	constant "scrabble/app/models/const"
	"time"
)

type PlayerInfo struct {
	Name              string
	ID                string
	VirtualPlayerType *constant.VirtualPlayerType
}

type Player struct {
	Name              string
	ID                string
	CompletedWords    []string
	Easel             []Letter
	IsPlaying         bool
	Score             int
	HasAbandon        bool
	IsVirtual         bool
	PrivateObjective  *Objective
	VirtualPlayerType *constant.VirtualPlayerType
}

type VirtualPlayerName struct {
	Name       string
	Type       constant.VirtualPlayerType
	IsReadonly bool
}

type GameHistory struct {
	Started       time.Time
	Duration      int
	Creator       Player
	CreatorScore  int
	OponentScore  int
	Opponent      Player
	Mode          constant.GameMode
	GameAbandoned bool
}
