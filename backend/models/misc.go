package models

type ErrorResponse struct {
	Status bool `json:"Status"`
	Message string `json:"Message"`
}
