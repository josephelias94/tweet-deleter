package constants

var (
	ERROR_AUTH_TOKEN_EXCHANGE            = SFX_AUTH + "Failed to convert an authorization code into a token | "
	ERROR_ENVS_LOADING_FILE              = SFX_ENVS + "Error loading .env file | "
	ERROR_ENVS_LOADING_VARIABLE          = SFX_ENVS + "Error loading variable. Check if .env file is loaded or the environment variable is not empty | "
	ERROR_ROUTES_SERVER                  = SFX_ROUTES + "Error in the server | "
	ERROR_TW_JSON_CONVERTING             = SFX_TW + "Error converting JSON | "
	ERROR_TW_JSON_INVALID                = SFX_TW + "Invalid JSON provided | "
	ERROR_TW_REQUEST_BUILDING            = SFX_TW + "Error building request | "
	ERROR_TW_REQUEST_DURING              = SFX_TW + "Error during request | "
	ERROR_TW_REQUEST_RESPONSE            = SFX_TW + "Error getting the response from request | "
	ERROR_TW_TWEETS_FAILED_STATUS_CODE   = SFX_TW_TWEETS + "Request to get tweets failed | "
	ERROR_TW_TWEETS_USER_UNSET           = SFX_TW_TWEETS + "User is not set"
	ERROR_TW_T_DELETE_FAILED_STATUS_CODE = SFX_TW_T_DELETE + "Request to delete a tweet failed | "
	ERROR_TW_USER_FAILED_STATUS_CODE     = SFX_TW_USER + "Request to get user failed | "
	ERROR_VALIDATOR_STRUCT               = SFX_VALIDATOR + "Failed struct validation | "
)
