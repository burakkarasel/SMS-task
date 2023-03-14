package models

// GenericResponse holds the data that returned from a controller
type GenericResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Messages   []string `json:"messages"`
	Data       any      `json:"data"`
}
