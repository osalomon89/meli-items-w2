package presenter

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
