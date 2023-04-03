package middleware

import (
	"net/http"
	"strings"

	"github.com/Calmantara/go-common/pkg/response"
	"github.com/Calmantara/go-dts-user/pkg/encoding"
	"github.com/gin-gonic/gin"
)

type HeaderKey string

func (h HeaderKey) String() string {
	return string(h)
}

const (
	Authorization HeaderKey = "Authorization"

	BasicAuth string = "Basic "
)

// middleware di GIN
// masuk kedalam tipe funcion handler
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// auth header
		header := ctx.GetHeader(Authorization.String())
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "token is not found",
			})
			return
		}
		// get token
		token := strings.Split(header, "Basic ")
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "token is not found",
			})
			return
		}

		// header token is found
		payload, err := encoding.DecodeBase64(ctx, token[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "invalid token",
			})
			return
		}

		// check payload
		splitted := strings.Split(payload, ":")
		if len(splitted) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "invalid token",
			})
			return
		}

		// set username and password
		userBasicMap := map[string]string{
			"username": splitted[0],
			"password": splitted[1],
		}
		ctx.Set("userBasic", userBasicMap)
		ctx.Next()
	}
}
