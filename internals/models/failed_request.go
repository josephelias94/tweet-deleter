package models

type FailedRequest struct {
	Detail string `json:"detail,omitempty" validate:"required"`
	Status int    `json:"status,omitempty" validate:"required"`
	Title  string `json:"title,omitempty" validate:"required"`
	Type   string `json:"type,omitempty" validate:"required"`
}
