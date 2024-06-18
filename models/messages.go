package models

type Message struct {
	Destination string `json:"destination" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Timestamp   string `json:"timestamp" validate:"required"`
}
