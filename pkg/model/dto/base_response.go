package dto

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BaseResponse(success bool, msg string, data interface{}) Response {
	return Response{
		Success: success,
		Message: msg,
		Data:    data,
	}
}
