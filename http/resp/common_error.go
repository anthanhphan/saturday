package resp

import "net/http"

func ErrInternalServer(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusInternalServerError,
		err,
		err.Error(),
	)
}

func ErrInvalidRequest(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusBadRequest,
		err,
		err.Error(),
	)
}

func ErrMissingTokenInHeader(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"missing token in header",
	)
}

func ErrInvalidTokenFormat(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"token is invalid format",
	)
}

func ErrInvalidTokenSignature(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"token is invalid signature",
	)
}
