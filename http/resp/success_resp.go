package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResp struct {
	StatusCode int         `json:"status_code" extensions:"x-order=1"`        // HTTP status code of the response
	Message    string      `json:"message,omitempty" extensions:"x-order=2"`  // Optional message describing the response
	Metadata   interface{} `json:"metadata,omitempty" extensions:"x-order=3"` // Optional additional data
}

// NewSuccessResp creates a new success response with a status code of 200 (OK).
// It includes an optional message and metadata.
//
// Parameters:
//   - msg: A string message describing the success response
//   - data: Additional metadata to include in the response (can be any type)
//
// Returns:
//   - *SuccessResp: A pointer to the newly created SuccessResp instance
//
// Examples:
//
//	resp := NewSuccessResp("operation successful", map[string]interface{}{"id": 123})
//	// resp.StatusCode == 200
//	// resp.Message == "operation successful"
//	// resp.Metadata == map[string]interface{}{"id": 123}
//
//	resp := NewSuccessResp("", nil)
//	// resp.StatusCode == 200
//	// resp.Message == ""
//	// resp.Metadata == nil
func NewSuccessResp(msg string, data interface{}) *SuccessResp {
	return &SuccessResp{
		StatusCode: http.StatusOK,
		Message:    msg,
		Metadata:   data,
	}
}

// ResponseSuccess writes a success response to the HTTP response writer using Gin's JSON serialization.
//
// Parameters:
//   - ctx: The Gin context for the HTTP request
//   - response: The success response to write
//
// Examples:
//
//	ctx := // obtain Gin context
//	resp := NewSuccessResp("operation successful", nil)
//	ResponseSuccess(ctx, resp)
//	// Writes a JSON response with status code 200 and message "operation successful"
func ResponseSuccess(ctx *gin.Context, response *SuccessResp) {
	ctx.JSON(response.StatusCode, response)
}
