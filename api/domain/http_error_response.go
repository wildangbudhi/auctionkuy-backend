package domain

// HTTPErrorResponse is object for HTTP Error Response
type HTTPErrorResponse struct {
	Error string `json:"error" binding:"required"`
}
