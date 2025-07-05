package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error      bool        `json:"error"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Errors     interface{} `json:"errors,omitempty"`
}

func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	resp := SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func WriteError(w http.ResponseWriter, status int, message string, errors interface{}) {
	resp := ErrorResponse{
		Error:      true,
		Message:    message,
		StatusCode: status,
		Errors:     errors,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
