// unit test for / endpoint

package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"netflix-watchlist/controller"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.HomePage)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "OK response is expected")
}

// unit test for /api/movies endpoint and GetAllMovies function asserts for json response and check out put is a valid json
func TestGetAllMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/movies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllMovies)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "OK response is expected")
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json", "application/json header is expected")
	assert.JSONEq(t, rr.Body.String(), `[{"Id":1,"Movie":"The Dark Knight","Rating":5},{"Id":2,"Movie":"The Matrix","Rating":4},{"Id":3,"Movie":"Inception","Rating":5},{"Id":4,"Movie":"Interstellar","Rating":4}]`, "Response body should match the expected json")
}

// unit test for /api/movies/{id} endpoint and GetMovie function asserts for json response and check out put is a valid json
func TestGetMovie(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie/{id}", controller.GetMovie)
	reqURL := "/api/movie/1"
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "OK response is expected")
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json", "application/json header is expected")
	assert.JSONEq(t, rr.Body.String(), `{"Id":1,"Movie":"The Dark Knight","Rating":5}`, "Response body should match the expected json")
}

// unit test for /api/movie endpoint and CreateMovie function asserts for json response and check out put is a valid json

func TestCreateMovie(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie", controller.CreateMovie)

	// Define the JSON payload for the request body
	requestBody := []byte(`{"Id": 5, "Movie": "KGF", "Rating": 100}`)

	// Create a new POST request with the JSON payload
	reqURL := "/api/movie"
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Perform assertions on the response
	assert.Equal(t, http.StatusOK, rr.Code, "OK response is expected")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "application/json header is expected")
	assert.JSONEq(t, `{"Id":5,"Movie":"KGF","Rating":100}`, rr.Body.String(), "Response body should match the expected json")
}

// unit test for /api/movie/{id} endpoint and DeleteMovie function asserts for json response and check out put is a valid json
func TestDeleteMovie(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie)
	reqURL := "/api/movie/5"
	req, err := http.NewRequest("DELETE", reqURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "OK response is expected")
	assert.Equal(t, rr.Body.String(), "Movie with ID = 5 was deleted.", "Response body should match the expected json")
}

// unit test for /api/movie/{id} endpoint and UpdateMovieRating function asserts for json response and check output

func TestUpdateMovieRating(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovieRating)

	// Define the JSON payload for the request body
	requestBody := []byte(`{"Id": 5, "Movie": "KGF", "Rating": 5}`)

	// Create a new PUT request with the JSON payload
	reqURL := "/api/movie/5"
	req, err := http.NewRequest("PUT", reqURL, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Serve the request and record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Perform assertions on the response
	assert.Equal(t, http.StatusOK, rr.Code, "OK response is expected")
	assert.Equal(t, rr.Body.String(), "Movie with ID = 5 was updated.", "Response body should match the expected json")
}
