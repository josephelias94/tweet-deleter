package models

type (
	User struct {
		Id       string `json:"id,omitempty" validate:"required"`
		Name     string `json:"name,omitempty" validate:"required"`
		Username string `json:"username,omitempty" validate:"required"`
	}

	GetUserResponse struct {
		Data User `json:"data,omitempty" validate:"required"`
	}
)
