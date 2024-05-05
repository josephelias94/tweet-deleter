package models

type (
	Errors struct {
		Message    string `json:"message,omitempty"`
		Parameters any    `json:"parameters,omitempty"`
	}

	FailedRequest struct {
		Detail string  `json:"detail,omitempty" validate:"required"`
		Errors *Errors `json:"errors,omitempty"`
		Status *int    `json:"status,omitempty"`
		Title  string  `json:"title,omitempty" validate:"required"`
		Type   string  `json:"type,omitempty" validate:"required"`
	}
)
