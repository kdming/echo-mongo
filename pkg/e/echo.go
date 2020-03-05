package e

// 自定义code
func ReturnBody(code int, data interface{}, msg string) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = msg
	response["data"] = data
	return response
}

// 正常返回
func SuccessReturn(data interface{}, msg string) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = 0
	response["msg"] = msg
	response["data"] = data
	return response
}

// 异常返回
func ErrorReturn(msg string) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = 1
	response["msg"] = msg
	response["data"] = nil
	return response
}
