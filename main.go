package main

import (
	"net/http"
	"netflix-watchlist/router"
)

func main() {
	//http listen and server
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
