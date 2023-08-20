package router

import (
	"netflix-watchlist/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.HomePage).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovieRating).Methods("PUT")
	router.HandleFunc("/api/directors", controller.GetAllDirectors).Methods("GET")
	router.HandleFunc("/api/director/{id}", controller.GetDirector).Methods("GET")
	router.HandleFunc("/api/director", controller.CreateDirector).Methods("POST")
	router.HandleFunc("/api/countries", controller.GetAllCountries).Methods("GET")
	return router
}
