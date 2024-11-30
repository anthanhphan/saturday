package metadata

import (
	"context"
	"fmt"

	"github.com/anthanhphan/saturday/http/constant/ctxkey"
	"github.com/anthanhphan/saturday/http/requester"
	"github.com/gin-gonic/gin"
)

func SetRequesterContextHeader(ctx *gin.Context) context.Context {
	if requesterVal, exists := ctx.Get(string(ctxkey.CtxRequesterKey)); exists {
		if r, ok := requesterVal.(requester.CtxRequester); ok {
			return context.WithValue(ctx.Request.Context(), ctxkey.CtxRequesterKey, r)
		}
	}

	return ctx.Request.Context()
}

func GetRequester(ctx context.Context) (requester.CtxRequester, error) {
	req, ok := ctx.Value(ctxkey.CtxRequesterKey).(requester.CtxRequester)
	if !ok {
		return nil, fmt.Errorf("requester not found in context")
	}
	return req, nil
}
