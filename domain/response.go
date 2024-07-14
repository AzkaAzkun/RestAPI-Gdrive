package domain

type Response struct {
	Success bool
	Error   string
	Data    any
}

type ServiceResponse struct {
	Code    int
	Success bool
	Error   string
	Data    any
}
