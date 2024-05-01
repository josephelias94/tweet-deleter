package models

type Username struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
}

type GetUsernameResponse struct {
	Data Username `json:"data,omitempty"`
}
