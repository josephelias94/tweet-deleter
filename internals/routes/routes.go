package routes

import (
	"log"
	"net/http"

	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/executioner"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauth2/callback", callback)

	log.Printf("%vPort: %v\n\n", constants.INFO_ROUTES_START_SERVER, constants.PORT)

	if err := http.ListenAndServe(constants.PORT, mux); err != nil {
		log.Fatalf("%vErrorMessage: \"%v\"\n\n", constants.ERROR_ROUTES_SERVER, err)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	executioner.StartDeletingStuff(code)
}
