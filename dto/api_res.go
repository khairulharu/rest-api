package dto

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
