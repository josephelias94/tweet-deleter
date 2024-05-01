package models

type Deleted struct {
	Deleted bool `json:"deleted,omitempty"`
}

type DeleteTwitterResponse struct {
	Data Deleted `json:"data,omitempty"`
}
