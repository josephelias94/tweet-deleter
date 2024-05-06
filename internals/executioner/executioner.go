package executioner

import (
	"fmt"
	"os"
	"time"

	"github.com/josephelias94/tweet-deleter/internals/authorizer"
	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
)

var (
	auth   *authorizer.Authorizer
	client twitter.Client
)

func StartDeletingStuff(code string) {
	auth = authorizer.GetInstance()
	auth.SetToken(code)

	client = twitter.Client{
		AuthorizedClient: auth.GetAuthorizedClient(),
	}

	client.SetUser("assimfalouojose")

	// deleteTweetsRecursively()
	deleteLikedTweetsRecursively()
}

func deleteTweetsRecursively() {
	counter := 0

	tweets := client.GetTweets()
	if len(tweets) == 0 {
		fmt.Print("Nothing to iterate over. Bye, bye")
		os.Exit(0)
	}

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

	time.Sleep(time.Duration(constants.RATE_LIMIT_IN_SECONDS) * time.Second)
	fmt.Printf("Iteration of %v tweet(s) finished. Starting over again \n\n", len(tweets))

	deleteTweetsRecursively()
}

func deleteLikedTweetsRecursively() {
	counter := 0

	tweets := client.GetLikedTweets()
	if len(tweets) == 0 {
		fmt.Print("Nothing to iterate over. Bye, bye")
		os.Exit(0)
	}

	for _, tweet := range tweets {
		if isMultipleOfFive(counter) {
			time.Sleep(time.Duration(constants.RATE_LIMIT_IN_SECONDS) * time.Second)
		}

		fmt.Printf(buildCounterMessage(counter, tweet.Id, len(tweets)))
		status, err := client.DeleteLikedTweet(tweet.Id)

		if err != nil {
			fmt.Printf("%v \n\n", err.Error())
		} else if status == false {
			fmt.Printf("Wasn't possible to delete a like | No errors available \n\n")
		} else if status == true {
			fmt.Printf("Deleted successfully \n\n")
		}

		counter += 1
	}

	time.Sleep(time.Duration(constants.RATE_LIMIT_IN_SECONDS) * time.Second)
	fmt.Printf("Iteration of %v like(s) finished. Starting over again \n\n", len(tweets))

	deleteLikedTweetsRecursively()
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
