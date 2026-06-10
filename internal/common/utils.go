package common

func ErrorResponse(data string) map[string]string {
	return map[string]string{
		"status": "error",
		"data":   data,
	}
}

func SuccessResponse(data any, message string) map[string]any {
	return map[string]any{
		"status":  "success",
		"data":    data,
		"message": message,
	}
}
