package models

type (
	Tweet struct {
		Id                  string   `json:"id,omitempty" validate:"required"`
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids,omitempty" validate:"required"`
		Text                string   `json:"text,omitempty" validate:"required"`
	}

	Meta struct {
		NewestId    string `json:"newest_id,omitempty" validate:"required"`
		NextToken   string `json:"next_token,omitempty" validate:"required"`
		OldestId    string `json:"oldest_id,omitempty" validate:"required"`
		ResultCount int    `json:"result_count,omitempty" validate:"required"`
	}

	Deleted struct {
		Deleted bool `json:"deleted,omitempty" validate:"required"`
	}

	GetTweetsResponse struct {
		Data []Tweet `json:"data,omitempty" validate:"required"`
		Meta Meta    `json:"meta,omitempty" validate:"required"`
	}

	DeleteTweetResponse struct {
		Data Deleted `json:"data,omitempty" validate:"required"`
	}
)
