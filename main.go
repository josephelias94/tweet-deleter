package main

import (
	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/envs"
	"github.com/josephelias94/tweet-deleter/internals/twitter"
	"golang.org/x/oauth2"
)

func main() {
	envs.Load()

	execTwitter()

	// routes.StartServer()
}

func execTwitter() {
	// fmt.Printf("MAIN | CLIENT_ID: %v\n", envs.GetClientId())
	// fmt.Printf("MAIN | CLIENT_SECRET: %v\n", envs.GetClientSecret())
	// fmt.Printf("MAIN | REDIRECT_URI: %v\n", envs.GetRedirectUri())
	// fmt.Printf("MAIN | AUTHORIZE: %v\n", constants.AUTHORIZE)
	// fmt.Printf("MAIN | OAUTH2_TOKEN: %v\n", constants.OAUTH2_TOKEN)
	// fmt.Printf("MAIN | TWITTER_SCOPES: %v\n", constants.TWITTER_SCOPES)

	client := twitter.Client{
		Config: oauth2.Config{
			ClientID:     envs.GetClientId(),
			ClientSecret: envs.GetClientSecret(),
			RedirectURL:  envs.GetRedirectUri(),
			Scopes:       constants.TWITTER_SCOPES,
			Endpoint: oauth2.Endpoint{
				AuthURL:  constants.AUTHORIZE,
				TokenURL: constants.OAUTH2_TOKEN,
			},
		},
	}

	client.Authorize()
	// client.SetUser("assimfalouojose")

	// fmt.Printf("MAIN | User: %v\n", client.User)
}
