package dto

// APIResponse API 的回傳值
type APIResponse struct {
	Meta APIResponseMeta `json:"meta"`
	Code string          `json:"code"`
	Message string          `json:"message"`
	Data    interface{}     `json:"data"`
}

// APIResponseMeta API Meta 的回傳值
type APIResponseMeta struct {
	RequestID string `json:"requestID"`
}
