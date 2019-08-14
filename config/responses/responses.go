package responses

//GeneralResponse returns a standard response format
type GeneralResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
}
