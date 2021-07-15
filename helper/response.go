package helper

type Response map[string]interface{}

func Success(data interface{}, error interface{}, message string) Response {
	response := Response{"status": "success", "errors": error, "data": data, "message": message}
	return response
}

func Failed(error interface{}, message string) Response {
	response := Response{"status": "success", "errors": error, "data": nil, "message": message}
	return response
}
