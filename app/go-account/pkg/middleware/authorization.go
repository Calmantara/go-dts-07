package middleware

import (
	"net/http"
	"strings"

	tokenmodel "github.com/Calmantara/go-account/modules/models/token"
	"github.com/Calmantara/go-account/pkg/crypto"
	"github.com/Calmantara/go-common/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	HeaderKey  string
	ContextKey string
)

func (h HeaderKey) String() string {
	return string(h)
}

func (h ContextKey) String() string {
	return string(h)
}

const (
	Authorization HeaderKey = "Authorization"

	AccessClaim ContextKey = "access_claim"

	BasicAuth  string = "Basic "
	BearerAuth string = "Bearer "
)

func BearerOAuth() gin.HandlerFunc {
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
		token := strings.Split(header, BearerAuth)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "token is not found",
			})
			return
		}

		// header token is found
		var claim tokenmodel.AccessClaim
		err := crypto.ParseJWT(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "invalid token",
			})
			return
		}
		ctx.Set(AccessClaim.String(), claim)
		ctx.Next()
	}
}

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
		token := strings.Split(header, BasicAuth)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Message: response.Unauthorized,
				Error:   "token is not found",
			})
			return
		}

		// header token is found
		payload, err := crypto.DecodeBase64(ctx, token[1])
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
