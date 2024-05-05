package constants

var (
	API_URL = "https://api.twitter.com/2"

	AUTHORIZE    = "https://twitter.com/i/oauth2/authorize"
	OAUTH2_TOKEN = API_URL + "/oauth2/token"

	DELETE_TWEET         = API_URL + "/tweets/:id"
	GET_TWEETS_BY_USER   = API_URL + "/users/:id/tweets?max_results=100" // 40
	GET_USER_BY_USERNAME = API_URL + "/users/by/username/:username"
)
