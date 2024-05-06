package models

type (
	DeleteLikedTweetData struct {
		Liked bool `json:"liked,omitempty" validate:"required"`
	}

	DeleteTweetData struct {
		Deleted bool `json:"deleted,omitempty" validate:"required"`
	}

	GetLikedTweetsMeta struct {
		NextToken   string `json:"next_token,omitempty" validate:"required"`
		ResultCount int    `json:"result_count,omitempty" validate:"required"`
	}

	GetTweetsMeta struct {
		NewestId    string `json:"newest_id,omitempty" validate:"required"`
		NextToken   string `json:"next_token,omitempty" validate:"required"`
		OldestId    string `json:"oldest_id,omitempty" validate:"required"`
		ResultCount int    `json:"result_count,omitempty" validate:"required"`
	}

	Tweet struct {
		Id                  string   `json:"id,omitempty" validate:"required"`
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids,omitempty" validate:"required"`
		Text                string   `json:"text,omitempty" validate:"required"`
	}

	GetTweetsResponse struct {
		Data []Tweet       `json:"data,omitempty" validate:"required"`
		Meta GetTweetsMeta `json:"meta,omitempty" validate:"required"`
	}

	GetLikedTweetsResponse struct {
		Data []Tweet            `json:"data,omitempty" validate:"required"`
		Meta GetLikedTweetsMeta `json:"meta,omitempty" validate:"required"`
	}

	DeleteTweetResponse struct {
		Data DeleteTweetData `json:"data,omitempty" validate:"required"`
	}

	DeleteLikedTweetResponse struct {
		Data DeleteLikedTweetData `json:"data,omitempty" validate:"required"`
	}
)
