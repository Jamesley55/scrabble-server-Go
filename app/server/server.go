package server

import (
	"log"
	"net/http"
	gameController "scrabble/app/controllers/game"
	"scrabble/app/models"
	gameServices "scrabble/app/services/game"

	mux "github.com/gorilla/mux"
)

func StartServer() {
	var games = []models.Game{}
	gameService := &gameServices.GameService{
		GameSessions: games,
	}
	router := mux.NewRouter()
	game := &gameController.GameController{
		GameService: gameService,
		Router:      router,
	}

	const port string = ":8080"
	router.HandleFunc("/posts", game.ConfigureRouter().ServeHTTP)
	log.Println("server listening on port", port)
	http.ListenAndServe(port, router)

	log.Fatalln()
}
