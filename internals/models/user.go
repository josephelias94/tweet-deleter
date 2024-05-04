package models

type (
	User struct {
		Id       string `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Username string `json:"username,omitempty"`
	}

	GetUserResponse struct {
		Data User `json:"data,omitempty"`
	}
)
