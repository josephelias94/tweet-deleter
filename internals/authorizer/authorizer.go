package authorizer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/josephelias94/tweet-deleter/internals/constants"
	"github.com/josephelias94/tweet-deleter/internals/envs"
	"golang.org/x/oauth2"
)

var (
	lock           = &sync.Mutex{}
	singleInstance *Authorizer
)

type Authorizer struct {
	Config           oauth2.Config
	authorizedClient *http.Client
	ctx              context.Context
	verifier         string
}

func GetInstance() *Authorizer {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &Authorizer{
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
		}
	}

	return singleInstance
}

func (a *Authorizer) Authorize() {
	a.ctx = context.Background()
	a.verifier = oauth2.GenerateVerifier()

	url := a.Config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(a.verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v\n\n", url)
}

func (a *Authorizer) SetToken(code string) {
	tok, err := a.Config.Exchange(a.ctx, code, oauth2.VerifierOption(a.verifier))
	if err != nil {
		log.Fatalf("%vErrorMessage: \"%v\"", constants.ERROR_AUTH_TOKEN_EXCHANGE, err)
	}

	a.authorizedClient = a.Config.Client(a.ctx, tok)
}

func (a *Authorizer) GetAuthorizedClient() *http.Client {
	return a.authorizedClient
}
