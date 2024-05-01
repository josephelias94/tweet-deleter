package main

import (
	"fmt"

	"github.com/josephelias94/tweet-deleter/internals/envs"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
)

func main() {
	envs.Load()
	// routes.StartServer()

	client := &twitter.Client{
		Token: envs.GetBearerToken(),
	}

	client.SetUser("assimfalouojose")

	fmt.Println(client.User)
}
