package gameServices

import (
	"errors"
	"fmt"
	models "scrabble/app/models"
	constant "scrabble/app/models/const"
	grid "scrabble/app/models/grid"
	utils "scrabble/app/utils"
	"time"
)

type GameService struct {
	GameSessions []models.Game
}

func (s *GameService) InitGame(gameOptions models.GameOptions) (*models.Game, error) {
	// letterReserve := grid.LETTERS
	// playerEasel := Container.Get(EaselService).GeneratePlayerLetters(letterReserve)

	// creator := Container.Get(PlayerService).AddCreator(gameOptions, playerEasel)
	// grid := Container.Get(GridService).LoadGrid()

	// game := s.CreateGame(gameOptions, creator, grid, letterReserve)

	// s.AddGameSession(game)
	// Container.Get(WebsocketService).IO.Emit("updateGamesAvailable")
	// return game
	fmt.Print("hello from init Game ")
	return nil, errors.New("opps error")
}

func (s *GameService) JoinGame(joinOpt models.JoinMultiplayerOption, isVirtualPlayer bool) (*models.Game, error) {
	game := s.GetGameByID(joinOpt.GameId)
	if game == nil {
		return nil, errors.New("herre here")
	}
	if game.Capacity < 2 && !isVirtualPlayer {
		// Container.Get(PlayerService).AddOpponent(joinOpt.PlayerInfo, game)
		// Container.Get(WebsocketService).IO.Emit("updateGamesAvailable")
	} else if game.Capacity < 2 && isVirtualPlayer {
		// Container.Get(PlayerService).AddOpponent(joinOpt.PlayerInfo, game, true)
		// Container.Get(WebsocketService).IO.Emit("updateGamesAvailable")
	}
	timeNow := time.Now()
	game.StartedTime = &timeNow
	return game, errors.New("opps error")
}

func (s *GameService) HandleMaxSkip(game *models.Game) {
	if game.SkipCounter == constant.MAX_SKIP_COUNT {
		s.DeleteGame(game.Id)
	}
}

func (s *GameService) ConvertToSolo(gameId, virtualPlayerName string, virtualPlayerType constant.VirtualPlayerType) {
	virtualPlayer := models.PlayerInfo{Name: virtualPlayerName, ID: utils.GenerateID(), VirtualPlayerType: &virtualPlayerType}
	soloPlayerOption := models.JoinMultiplayerOption{PlayerInfo: virtualPlayer, GameId: gameId}
	s.JoinGame(soloPlayerOption, true)
	s.StartGame(gameId)
}

func (s *GameService) StartGame(gameID string) error {
	game := s.GetGameByID(gameID)
	if s.IsDictionnaryDeleted(gameID) {
		return errors.New("Dictionnary Deleted")
	}
	if game.Mode == "GameModeLog2990" {
		// Container.Get(ObjectivesService).AddObjectives(game)
	}
	// Container.Get(PlayerService).SetRandomStartingPlayer(game)
	// Container.Get(WebsocketService).IO.To(gameID).Emit("startGame", game)
	// Container.Get(TimerService).StartTimer(game)
	if game.Opponent.IsVirtual {
		// go Container.Get(VirtualPlayerService).StartVirtualPlayer(game)
	}
	return nil
}

func (s *GameService) GetGameByID(gameID string) *models.Game {
	for _, game := range s.GameSessions {
		if game.Id == gameID {
			return &game
		}
	}
	return nil
}

func (s *GameService) CreateGame(gameOptions models.GameOptions, creator *models.Player, grid []grid.Node, letterReserve []string) models.Game {
	timeNow := time.Now()
	return models.Game{
		Id:            gameOptions.PlayerId,
		Creator:       *creator,
		Opponent:      nil,
		Grid:          grid,
		LetterReserve: letterReserve,
		Mode:          string(gameOptions.GameMode),
		Time:          gameOptions.Time,
		Capacity:      1,
		SkipCounter:   0,
		StartedTime:   &timeNow,
	}
}

func (s *GameService) AddGameSession(game models.Game) {
	s.GameSessions = append(s.GameSessions, game)
}

func (s *GameService) DeleteGame(gameID string) {
	for i, game := range s.GameSessions {
		if game.Id == gameID {
			s.GameSessions = append(s.GameSessions[:i], s.GameSessions[i+1:]...)
			break
		}
	}
	// Container.Get(WebsocketService).IO.Emit("updateGamesAvailable")
}

func (s *GameService) IsDictionnaryDeleted(gameID string) bool {
	game := s.GetGameByID(gameID)
	if game == nil {
		return true
	}
	return false
	// return Container.Get(DictionaryService).IsDeleted(game.Mode)
}

func (s *GameService) UpdateGameHistory(gameID string) {
	game := s.GetGameByID(gameID)
	if game == nil {
		return
	}
	// Container.Get(GameHistoryService).UpdateGameHistory(game)
}

func (s *GameService) GetGameSessions() []models.Game {
	return s.GameSessions
}
