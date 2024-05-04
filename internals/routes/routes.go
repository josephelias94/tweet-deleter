package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josephelias94/tweet-deleter/internals/authorizer"
	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauth2/callback", callback)

	log.Printf("%vPort: %v", constants.INFO_ROUTES_START_SERVER, constants.PORT)

	if err := http.ListenAndServe(constants.PORT, mux); err != nil {
		log.Fatalf("%vErrorMessage: \"%v\"", constants.ERROR_ROUTES_SERVER, err)
	}
}

func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	auth := authorizer.GetInstance()

	auth.SetToken(code)
	execTwitterThings(auth)

	w.Write([]byte("Welcome to /callback endpoint | code: " + code))
}

func execTwitterThings(auth *authorizer.Authorizer) {
	client := twitter.Client{
		AuthorizedClient: auth.GetAuthorizedClient(),
	}

	client.SetUser("assimfalouojose")
	tweets := client.GetTweets()

	for _, tweet := range tweets {
		fmt.Println("Attempting to delete tweet id " + tweet.Id)

		status, err := client.DeleteTweet(tweet.Id)
		if err != nil {
			fmt.Println("Wasn't possible to delete tweet id " + tweet.Id + " | Error: " + err.Error() + "\n")

			return
		}

		if status == false {
			fmt.Println("Wasn't possible to delete tweet id " + tweet.Id + " | No errors available" + "\n")

			return
		}

		fmt.Println("Tweet id " + tweet.Id + " deleted successfully" + "\n")
	}
}
