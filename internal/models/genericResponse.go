package models

import "strings"

// GenericResponse holds the data that returned from a controller
type GenericResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Messages   []string `json:"messages"`
	Data       any      `json:"data"`
}

// CreateGenericResponse creates a new instance of GenericResponse
func CreateGenericResponse(succes bool, statusCode int, messages string, data any) GenericResponse {
	slc := strings.Split(messages, "\n")
	return GenericResponse{
		Success:    succes,
		StatusCode: statusCode,
		Messages:   slc,
		Data:       data,
	}
}
