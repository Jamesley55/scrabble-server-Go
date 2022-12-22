package gameController

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "scrabble/app/models"
	gameServices "scrabble/app/services/game"

	"github.com/gorilla/mux"
)

type GameController struct {
	GameService *gameServices.GameService
	// playerService *player.Service
	Router *mux.Router
}

func (gc *GameController) ConfigureRouter() *mux.Router {

	gc.Router.HandleFunc("/init/multi", func(w http.ResponseWriter, r *http.Request) {
		gameOption := models.GameOptions{}
		fmt.Println("here")

		if err := json.NewDecoder(r.Body).Decode(gameOption); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		game, err := gc.GameService.InitGame(gameOption)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(game)
	}).Methods(http.MethodPost)

	// gc.router.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
	// 	game, err := gc.gameService.JoinGame(r.Body)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if game == nil {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(game)
	// }).Methods(http.MethodPost)

	// gc.router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
	// 	var req struct {
	// 		GameID string `json:"gameId"`
	// 	}
	// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	if err := gc.gameService.StartGame(GameID); err != nil {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusCreated)
	// }).Methods(http.MethodPost)
	return gc.Router
}
