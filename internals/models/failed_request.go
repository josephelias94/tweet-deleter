package models

type FailedRequest struct {
	Detail string `json:"detail,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
}
