package common

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(data string) map[string]string {
	return map[string]string{
		"status": "error",
		"data":   data,
	}
}

func SendCreated(w *http.ResponseWriter, data any) {
	(*w).WriteHeader(http.StatusCreated)
	json.NewEncoder(*w).Encode(data)
}

func SendError(w *http.ResponseWriter, statusCode int, msg string) {
	(*w).Header().Add("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	json.NewEncoder(*w).Encode(ErrorResponse(msg))
}

func SendSuccess(w *http.ResponseWriter, data any) {
	(*w).Header().Add("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(SuccessResponse(data, ""))
}

func SuccessResponse(data any, message string) map[string]any {
	return map[string]any{
		"status":  "success",
		"data":    data,
		"message": message,
	}
}
