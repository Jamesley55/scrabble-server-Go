package gameController

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "scrabble/app/models"
	gameServices "scrabble/app/services/game"
	"scrabble/app/utils"

	"github.com/gorilla/mux"
)

type GameController struct {
	GameService *gameServices.GameService
	// playerService *player.Service
	Router *mux.Router
}

func (gc *GameController) ConfigureRouter() *mux.Router {
	route := gc.Router
	route.HandleFunc("/helloWorld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("<h1>Hello<h1>"))
	}).Methods(http.MethodGet)

	route.HandleFunc("/init/multi", func(w http.ResponseWriter, r *http.Request) {
		gameOption := models.GameOptions{}
		if err := json.NewDecoder(r.Body).Decode(&gameOption); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		utils.Info(gameOption)
		game, err := gc.GameService.InitGame(gameOption)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(game)
	}).Methods(http.MethodPost)

	route.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		multiplayerOption := models.JoinMultiplayerOption{}
		if err := json.NewDecoder(r.Body).Decode(&multiplayerOption); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		game, err := gc.GameService.JoinGame(multiplayerOption, false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if game == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(game)
	}).Methods(http.MethodPost)

	route.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			GameID string `json:"gameId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := gc.GameService.StartGame(req.GameID); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}).Methods(http.MethodPost)

	route.HandleFunc("/rejectPlayer", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			GameID string `json:"gameId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		game := gc.GameService.GetGameByID(request.GameID)
		if game == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// gc.playerService.RejectPlayer(game)
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodDelete)

	route.HandleFunc("/gameSession", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(gc.GameService.GameSessions)
	}).Methods(http.MethodGet)

	route.HandleFunc("/gameSession", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			GameID string `json:"gameId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// if err := gc.GameService.RemoveGameSession(request.GameID); err != nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		// }
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodDelete)

	route.HandleFunc("/playerCheck", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			GameID   string `json:"gameId"`
			PlayerID string `json:"playerId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		game := gc.GameService.GetGameByID(request.GameID)
		if game == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(false)
			return
		}
		// player := gc.layerService.GetPlayerByID(request.PlayerID, game)
		// if player == nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	json.NewEncoder(w).Encode(false)
		// 	return
		// }
		// if player.HasAbandon {
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(false)
		// 	return
		// }
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	}).Methods(http.MethodPost)

	route.HandleFunc("/reconnect", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			OldID  string `json:"oldId"`
			NewID  string `json:"newId"`
			GameID string `json:"gameId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// game := gc.GameService.ReconnectToGame(request.OldID, request.NewID, request.GameID)
		// if game == nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	json.NewEncoder(w).Encode("Reconnection timeout")
		// 	return
		// }
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(game)
	}).Methods(http.MethodPost)
	route.HandleFunc("/convertSolo", func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			GameID            string `json:"gameId"`
			OpponentName      string `json:"oponentName"`
			VirtualPlayerType string `json:"virtualPlayerType"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// if err := gc.GameService.ConvertToSolo(request.GameID, request.OpponentName, request.VirtualPlayerType); err != nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	return
		// }
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)

	return gc.Router
}
