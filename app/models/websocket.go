package models

type ClientEvents struct {
	Message            func(msg string, room string)
	Command            func(command Command, room string)
	JoinRoom           func(room string)
	LeaveRoom          func(room string)
	RemoveOpponent     func(room string)
	DeleteGame         func(room string)
	AbandonGame        func(room string)
	ValidateName       func(joinMultiplayerOption JoinMultiplayerOption)
	UpdateName         func(playerName string)
	UpdateScore        func(score int, gameId string)
	UpdateEasel        func(easel []Letter, playerId string, gameId string)
	ShadowPlaceLetters func(placeCommand PlaceCommand, game Game, playerId string)
	RequestGameUpdate  func(gameId string)
}

type ServerEvents struct {
	Message                 func(msg ChatMessage)
	JoinRoom                func(room string)
	LeaveRoom               func(room string)
	PlayerJoined            func(player PlayerInfo)
	PlayerInfo              func(player PlayerInfo)
	PlayerLeft              func(player PlayerInfo)
	PlayerAbandoned         func()
	StartGame               func(game Game)
	Rejected                func(player PlayerInfo)
	UpdateGamesAvailable    func()
	ValidName               func(isValid bool)
	NameUpdated             func()
	CommandSuccess          func(success bool, command Command)
	UpdateGame              func(game Game)
	UpdateTimer             func(timer int)
	EndOfGame               func()
	ShadowPlaceLetters      func(command PlaceCommand)
	BadPlaceCommandPosition func()
	VirtualPlayerChange     func(virtualPlayers []VirtualPlayerName)
	ResetHistory            func(history []GameHistory)
	NewDict                 func(dict []DictHeaders)
	NonExistingDict         func()
}

type SocketInfo struct {
	GameId string
	Player Player
}

type SocketAuth struct {
	PlayerName     string
	LastPlayerInfo *PlayerInfo
}
