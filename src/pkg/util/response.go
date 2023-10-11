package util

import "gin-user-tasks/src/pkg/config"

type Response struct {
	App          string      `json:"app"`
	RequestID    string      `json:"request_id,omitempty"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		App:     config.Env.Name,
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
