package main

import (
	"net/http"
	"webapp/internal/config"
)

func main() {

	// load config

	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

	})

	//

}
