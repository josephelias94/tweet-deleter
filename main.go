package main

import (
	"github.com/josephelias94/tweet-deleter/internals/authorizer"
	"github.com/josephelias94/tweet-deleter/internals/envs"
	"github.com/josephelias94/tweet-deleter/internals/routes"
)

func main() {
	envs.Load()

	auth := authorizer.GetInstance()
	go auth.Authorize()

	routes.StartServer()
}
