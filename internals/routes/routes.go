package routes

import (
	"log"
	"net/http"

	"github.com/josephelias94/tweet-deleter/internals/constants"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauth2/callback", callback)

	log.Printf("Starting server on port %v", constants.PORT)

	if err := http.ListenAndServe(constants.PORT, mux); err != nil {
		log.Fatalf("routes: %v", err)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to /callback endpoint"))
}
