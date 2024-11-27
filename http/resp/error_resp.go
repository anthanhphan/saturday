package resp

import (
	"errors"
)

type ErrorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Log        string `json:"-"`
	RootErr    error  `json:"-"`
}

// NewErrorResp creates a new custom error response.
// It constructs an ErrorResp with a status code, message, and optional root error.
//
// Parameters:
//   - statusCode: The HTTP status code for the error response
//   - root: The root error that caused this error (can be nil)
//   - msg: The error message to be displayed
//
// Returns:
//   - *ErrorResp: A pointer to the newly created ErrorResp instance
//
// Examples:
//
//	err := errors.New("database connection failed")
//	resp := NewErrorResp(500, err, "internal server error")
//	// resp.StatusCode == 500
//	// resp.Message == "internal server error"
//	// resp.Log == "database connection failed"
//
//	resp := NewErrorResp(400, nil, "invalid input")
//	// resp.StatusCode == 400
//	// resp.Message == "invalid input"
//	// resp.Log == "invalid input"
func NewErrorResp(statusCode int, root error, msg string) *ErrorResp {
	if root != nil {
		return &ErrorResp{
			StatusCode: statusCode,
			RootErr:    root,
			Message:    msg,
			Log:        root.Error(),
		}
	}

	return &ErrorResp{
		StatusCode: statusCode,
		RootErr:    errors.New(msg),
		Message:    msg,
		Log:        msg,
	}
}

// RootError returns the underlying root cause of the error chain.
// It traverses the error chain to find the original error.
//
// Returns:
//   - error: The root error in the error chain
//
// Examples:
//
//	rootErr := errors.New("original error")
//	resp1 := NewErrorResp(500, rootErr, "first wrapper")
//	resp2 := NewErrorResp(500, resp1, "second wrapper")
//
//	err := resp2.RootError()
//	// err.Error() == "original error"
func (e *ErrorResp) RootError() error {
	var err *ErrorResp

	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}

	return e.RootErr
}

// Error implements the error interface and returns the error message.
// It provides the error message of the root error.
//
// Returns:
//   - string: The error message of the root error
//
// Examples:
//
//	resp := NewErrorResp(500, errors.New("db error"), "server error")
//	err := resp.Error()
//	// err == "db error"
func (e *ErrorResp) Error() string {
	return e.RootError().Error()
}
