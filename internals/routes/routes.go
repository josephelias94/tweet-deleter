package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/josephelias94/tweet-deleter/internals/authorizer"
	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
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
	auth := authorizer.GetInstance()

	auth.SetToken(code)
	execTwitterThings(auth)

	w.Write([]byte("Welcome to /callback endpoint | code: " + code))
}

func execTwitterThings(auth *authorizer.Authorizer) {
	counter := 0

	client := twitter.Client{
		AuthorizedClient: auth.GetAuthorizedClient(),
	}

	client.SetUser("assimfalouojose")
	tweets := client.GetTweets()

	for _, tweet := range tweets {

		if isMultipleOfFive(counter) {
			time.Sleep(time.Duration(constants.RATE_LIMIT_IN_SECONDS) * time.Second)
		}

		fmt.Printf(buildCounterMessage(counter, tweet.Id, len(tweets)))
		status, err := client.DeleteTweet(tweet.Id)

		if err != nil {
			fmt.Printf("%v \n\n", err.Error())
		} else if status == false {
			fmt.Printf("Wasn't possible to delete tweet | No errors available \n\n")
		} else if status == true {
			fmt.Printf("Deleted successfully \n\n")
		}

		counter += 1
	}
}

func isMultipleOfFive(number int) bool {
	if number == 0 {
		return false // skip it the first iteration
	}

	return number%5 == 0
}

func buildCounterMessage(counter int, id string, length int) string {
	return fmt.Sprintf("%v/%v | Attempting to delete tweet id: %v\n", counter+1, length, id)
}
