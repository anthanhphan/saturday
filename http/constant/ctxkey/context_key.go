package ctxkey

type ctxKeyType string

const (
	CtxRequestIdKey ctxKeyType = "X-Request-ID"
	CtxRequesterKey ctxKeyType = "CONTEXT_REQUESTER"
)
