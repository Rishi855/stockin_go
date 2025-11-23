package models

type CustomResponse struct {
	StatusCode int         `json:"status"`
	IsSuccess  bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}