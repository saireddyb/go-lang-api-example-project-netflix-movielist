package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"netflix-watchlist/model"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []model.MoviesList

// fake data
func init() {
	movies = []model.MoviesList{
		{Id: 1, Movie: "The Dark Knight", Rating: 5},
		{Id: 2, Movie: "The Matrix", Rating: 4},
		{Id: 3, Movie: "Inception", Rating: 5},
		{Id: 4, Movie: "Interstellar", Rating: 4},
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllMovies() []model.MoviesList {
	return movies
}
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	result := getAllMovies()
	json.NewEncoder(w).Encode(result)
}

func getMovie(id int) model.MoviesList {
	for _, movie := range movies {
		if movie.Id == id {
			return movie
		}
	}
	return model.MoviesList{}
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)
	// convert the id from string to int
	i, _ := strconv.Atoi(params["id"])
	result := getMovie(i)
	json.NewEncoder(w).Encode(result)
}

func createMovie(movie model.MoviesList) {
	movies = append(movies, movie)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.MoviesList
	_ = json.NewDecoder(r.Body).Decode(&movie)
	createMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(id int) {
	for index, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	deleteMovie(i)
	fmt.Fprint(w, "Movie with ID = "+params["id"]+" was deleted.")
}

func updateMovie(id int, movie model.MoviesList) {
	for index, item := range movies {
		if item.Id == id {
			movies[index].Rating = movie.Rating
		}
	}
}

func UpdateMovieRating(w http.ResponseWriter, r *http.Request) {
	var movie model.MoviesList
	_ = json.NewDecoder(r.Body).Decode(&movie)
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	updateMovie(i, movie)
	fmt.Fprint(w, "Movie with ID = "+params["id"]+" was updated.")
}
