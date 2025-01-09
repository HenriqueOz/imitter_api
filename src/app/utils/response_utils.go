package utils

func ResponseError(err error, details any) map[string]any {
	data := map[string]any{
		"error":   err.Error(),
		"details": details,
	}

	return data
}

func ResponseSuccess(data any) map[string]any {
	return map[string]any{
		"data": data,
	}
}
