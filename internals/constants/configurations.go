package constants

var (
	PORT                  = ":4000"
	TWITTER_SCOPES        = []string{"tweet.read", "tweet.write", "users.read", "offline.access"}
	RATE_LIMIT_IN_SECONDS = 960 // 16 minutes
)
