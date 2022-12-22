package playerServices

/*
import "math/rand"

type PlayerService struct{}

func (s *PlayerService) AddCreator(gameOptions GameOptions, playerEasel []Letter) Player {
	return Player{
		ID:           gameOptions.PlayerID,
		Name:         gameOptions.PlayerName,
		IsPlaying:    true,
		Completed:    []string{},
		Easel:        playerEasel,
		Score:        0,
		HasAbandoned: false,
		IsVirtual:    false,
	}
}

func (s *PlayerService) AddOpponent(playerInfo PlayerInfo, game *Game, isVirtual bool) {
	playerEasel := Container.Get(EaselService).GeneratePlayerLetters(game.LetterReserve)
	opponent := Player{
		ID:           playerInfo.ID,
		Name:         playerInfo.Name,
		IsPlaying:    false,
		Completed:    []string{},
		Easel:        playerEasel,
		Score:        0,
		HasAbandoned: false,
		IsVirtual:    isVirtual,
	}
	if playerInfo.VirtualPlayerType != "" {
		opponent.VirtualPlayerType = playerInfo.VirtualPlayerType
	}
	game.Opponent = &opponent
	game.Capacity++
}

func (s *PlayerService) RejectPlayer(game *Game) {
	player := PlayerInfo{
		ID:   game.Creator.ID,
		Name: game.Creator.Name,
	}
	s.RemoveOpponent(game)
	Container.Get(WebsocketService).IO.To(game.ID).Emit("rejected", player)
}

func (s *PlayerService) RemoveOpponent(game *Game) {
	if game.Opponent == nil {
		return
	}
	game.Opponent = nil
	game.Capacity--
	Container.Get(WebsocketService).IO.Emit("updateGamesAvailable")
}

func (s *PlayerService) GetPlayerByID(playerID string, game *Game) *Player {
	if game.Creator.ID == playerID {
		return game.Creator
	} else if game.Opponent != nil && game.Opponent.ID == playerID {
		return game.Opponent
	} else {
		return nil
	}
}

func (s *PlayerService) IncrementPlayerScore(score int, playerID string, game *Game) {
	player := s.GetPlayerByID(playerID, game)
	if player == nil {
		return
	}
	player.Score += score
}

func (s *PlayerService) SwitchPlayerTurn(game *Game) {
	if game.Opponent != nil && game.Opponent.IsPlaying {
		game.Creator.IsPlaying = true
		game.Opponent.IsPlaying = false
	} else if game.Opponent != nil && game.Creator.IsPlaying {
		game.Creator.IsPlaying = false
		game.Opponent.IsPlaying = true
	}
	game.Timer =
		game.Time
	Container.Get(GameService).HandleMaxSkip(game)
}

func (s *PlayerService) GetRemainingLettersScore(player *Player) int {
	score := 0
	for _, letter := range player.Easel {
		score += letter.Value
	}
	return score
}

func (s *PlayerService) UpdateEasel(easel []Letter, playerID string, game *Game) {
	player := s.GetPlayerByID(playerID, game)
	if player == nil {
		return
	}
	player.Easel = easel
}

func (s *PlayerService) SetRandomStartingPlayer(game *Game) {
	if rand.Intn(2) == 1 {
		s.SwitchPlayerTurn(game)
	}
}

func (s *PlayerService) UpdatePlayerFinalScore(player *Player) {
	player.Score -= s.GetRemainingLettersScore(player)
	if player.Score < 0 {
		player.Score = 0
	}
}
*/
