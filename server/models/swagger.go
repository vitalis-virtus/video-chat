package models

type HTTPSuccess struct {
	Success bool `json:"success" example:"true"`
}

type HTTPError struct {
	Error string `json:"error"`
}

type CreateChannelRes struct {
	ID string `json:"id" example:"66cd5b0d-4f91-43a6-96b0-f2ae9e3863d1"`
}
