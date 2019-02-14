package entity

// JSONResponse _
type JSONResponse struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"http_status"`
}

// NewJSONResponse _
func NewJSONResponse(message string, httpStatus int) JSONResponse {
	return JSONResponse{
		Message:    message,
		HTTPStatus: httpStatus,
	}
}
