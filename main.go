package main

import (
	"fmt"

	"github.com/josephelias94/tweet-deleter/internals/envs"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
)

func main() {
	envs.Load()

	execTwitter()

	// routes.StartServer()
}

func execTwitter() {
	client := twitter.Client{
		Token: envs.GetBearerToken(),
	}

	client.SetUser("assimfalouojose")
	fmt.Printf("user: %v\n", client.User)

	tweets := client.GetTweets()
	fmt.Printf("tweet: %v\n", tweets[9])
}
