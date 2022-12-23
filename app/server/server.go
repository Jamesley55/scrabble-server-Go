package server

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	gameController "scrabble/app/controllers/game"
	"scrabble/app/models"
	gameServices "scrabble/app/services/game"

	socket "scrabble/app/services/socket"

	"github.com/MadAppGang/httplog"
	mux "github.com/gorilla/mux"
)

func StartServer() {

	router := mux.NewRouter()
	const port string = ":8080"
	config(router)
	socket.NewWebsocketService()
	bindRoute(router)
	router.HandleFunc("/api/dictionaries", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
	}).Methods(http.MethodGet)
	log.Println("server listening on port", port)
	http.ListenAndServe(port, router)

	log.Fatalln()
}

func config(router *mux.Router) {
	// Middlewares configuration
	router.Use(httplog.Logger)
	router.Use(headerMiddleware)
	router.Use(panicRecovery)
	router.Use(CORSMiddleware)

}
func bindRoute(router *mux.Router) {

	game := &gameController.GameController{
		GameService: &gameServices.GameService{
			GameSessions: []models.Game{},
		},
		Router: router.PathPrefix("/api/game").Subrouter(),
	}

	router.HandleFunc("/", game.ConfigureRouter().ServeHTTP)
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func panicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
