package models

type Tweet struct {
	Id                  string   `json:"id,omitempty"`
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids,omitempty"`
	Text                string   `json:"text,omitempty"`
}

type Meta struct {
	NewestId    string `json:"newest_id,omitempty"`
	NextToken   string `json:"next_token,omitempty"`
	OldestId    string `json:"oldest_id,omitempty"`
	ResultCount int    `json:"result_count,omitempty"`
}

type Deleted struct {
	Deleted bool `json:"deleted,omitempty"`
}

type GetTweetResponse struct {
	Data []Tweet `json:"data,omitempty"`
	Meta Meta    `json:"meta,omitempty"`
}

type DeleteTweetResponse struct {
	Data Deleted `json:"data,omitempty"`
}
