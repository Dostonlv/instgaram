package models

type DefaultError struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
