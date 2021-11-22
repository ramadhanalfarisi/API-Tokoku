package helper

type Response map[string]interface{}

func Success(data interface{}, error interface{}, message string) Response {
	var response Response
	if data != nil{
		response = Response{"status": "success", "errors": error, "data": data, "message": message}
	}else{
		response = Response{"status": "success", "errors": error, "message": message}
	}
	return response
}

func FailedValidate(error interface{}, message []string) Response {
	response := Response{"status": "failed", "errors": error, "message": message}
	return response
}

func Failed(error interface{}, message string) Response {
	response := Response{"status": "failed", "errors": error, "message": message}
	return response
}
