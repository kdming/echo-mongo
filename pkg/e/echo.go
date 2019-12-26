package e

func ReturnBody(code int, data interface{}, msg string) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = msg
	response["data"] = data
	return response
}
