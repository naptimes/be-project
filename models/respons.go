package models

type Respon struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    Dashboard `json:"data"`
}
