package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"netflix-watchlist/model"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []model.Movies
var directors []model.Director
var countries []model.Nationality

// fake data
func init() {
	countries = []model.Nationality{
		{Country: "American"},
		{Country: "British"},
		{Country: "Australian"},
		{Country: "Indian"},
	}
	directors = []model.Director{
		{Id: 1, Name: "Christopher Nolan", Nationality: countries[1]},
		{Id: 2, Name: "Lana Wachowski", Nationality: countries[0]},
		{Id: 3, Name: "Raja mauli", Nationality: countries[3]},
	}
	movies = []model.Movies{
		{Id: 1, Name: "The Dark Knight", Rating: 5, Director: &directors[0]},
		{Id: 2, Name: "The Matrix", Rating: 4, Director: &directors[1]},
		{Id: 3, Name: "Inception", Rating: 5, Director: &directors[0]},
		{Id: 4, Name: "Interstellar", Rating: 4, Director: &directors[0]},
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllMovies() []model.Movies {
	return movies
}
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	result := getAllMovies()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getMovie(id int) model.Movies {
	for _, movie := range movies {
		if movie.Id == id {
			return movie
		}
	}
	return model.Movies{}
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	params := mux.Vars(r)
	// convert the id from string to int
	i, _ := strconv.Atoi(params["id"])
	result := getMovie(i)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func createMovie(movie model.Movies) {
	movies = append(movies, movie)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	directorID := movie.Director.Id
	var selectedDirector model.Director
	for _, dir := range directors {
		if dir.Id == directorID {
			selectedDirector = dir
			break
		}
	}
	if selectedDirector == (model.Director{Id: 0, Name: "", Nationality: model.Nationality{}}) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Director with ID %d not found", directorID)
		return
	}
	movie.Director = &selectedDirector
	createMovie(movie)
	w.Header().Set("Content-Type", "application/json")
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

func updateMovie(id int, movie model.Movies) {
	for index, item := range movies {
		if item.Id == id {
			movies[index].Rating = movie.Rating
		}
	}
}

func UpdateMovieRating(w http.ResponseWriter, r *http.Request) {
	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	updateMovie(i, movie)
	fmt.Fprint(w, "Movie with ID = "+params["id"]+" was updated.")
}

func getAllDirectors() []model.Director {
	return directors
}

func GetAllDirectors(w http.ResponseWriter, r *http.Request) {
	result := getAllDirectors()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getDirector(id int) model.Director {
	for _, dir := range directors {
		if dir.Id == id {
			return dir
		}
	}
	return model.Director{}
}

func GetDirector(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	result := getDirector(i)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func createDirector(director model.Director) {
	directors = append(directors, director)
}

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	var director model.Director
	_ = json.NewDecoder(r.Body).Decode(&director)
	createDirector(director)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}

func getAllCountries() []model.Nationality {
	return countries
}

func GetAllCountries(w http.ResponseWriter, r *http.Request) {
	result := getAllCountries()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
