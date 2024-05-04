package routes

import (
	"log"
	"net/http"

	"github.com/josephelias94/tweet-deleter/internals/constants"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauth2/callback", callback)

	log.Printf("%v Port: %v", constants.INFO_ROUTES_START_SERVER, constants.PORT)

	if err := http.ListenAndServe(constants.PORT, mux); err != nil {
		log.Fatalf("%v Error: \"%v\"", constants.ERROR_ROUTES_SERVER, err)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to /callback endpoint"))
}
