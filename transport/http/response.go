package http

import (
"net/http"
)

// Response struct
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  string                 `json:"errors,omitempty"`
}

// SetErrorResponse builder for error response
func (r *Response) SetErrorResponse(code int, message string) {
	r.Code = code
	r.Message = http.StatusText(code)
	r.Errors = message
}

// SetSuccessResponse builder for error response
func (r *Response) SetSuccessResponse(code int, result map[string]interface{}) {
	r.Code = code
	r.Message = http.StatusText(code)
	r.Data = result
}


