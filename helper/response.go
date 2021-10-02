package helper

import (
	"strings"
)

type Response struct {
	Status  bool 		`json:"status"`
	Message string 		`json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct {}

func BuildResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return response
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	response := Response{
		Status: false,
		Message: message,
		Errors: splittedError,
		Data: data,
	}
	return response
}