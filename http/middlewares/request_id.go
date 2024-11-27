package middlewares

import (
	"context"

	"github.com/anthanhphan/saturday/http/constant/ctxkey"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestId creates a middleware that generates and adds a unique request ID to each request's context.
// The request ID is generated using UUID v4 and stored in the context using ctxkey.CtxRequestIdKey.
//
// Returns:
//   - gin.HandlerFunc: Middleware function that adds request ID to context
//
// Examples:
//
//	router := gin.New()
//	router.Use(RequestId())
//	// Each request will now have a unique ID accessible via context
//	// Access ID in handlers: ctx.Request.Context().Value(ctxkey.CtxRequestIdKey).(string)
func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ctxkey.CtxRequestIdKey, uuid.New().String()))
		ctx.Next()
	}
}
