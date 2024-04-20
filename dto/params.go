package dto

type ResponseParam struct {
	Status     bool
	StatusCode int
	Message    string
	Data       any
}
