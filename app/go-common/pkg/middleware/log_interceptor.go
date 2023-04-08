package middleware

import (
	"github.com/Calmantara/go-common/pkg/context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CorrelationIDInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// try get from header
		val := ctx.GetHeader(context.CorrID.String())
		if val == "" {
			val = uuid.New().String()
		}

		ctx.Set(context.CorrID.String(), val)
		ctx.Next()
	}
}
