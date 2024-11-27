package middlewares

import (
	"net/http"

	"github.com/anthanhphan/saturday/http/resp"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recover() gin.HandlerFunc {
	log := zap.L().With(zap.String("prefix", "recover")).Sugar()
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)

				ctx.Header("Content-Type", "application/json")

				if commonErr, ok := err.(*resp.ErrorResp); ok {
					ctx.AbortWithStatusJSON(commonErr.StatusCode, commonErr)
					return
				}

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp.ErrInternalServer(err.(error)))
			}
		}()

		ctx.Next()
	}
}
