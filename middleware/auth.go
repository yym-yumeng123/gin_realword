package middleware

import (
	"gin_realword/logger"
	"gin_realword/security"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.New(ctx)

		token := ctx.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		claims, ok, err := security.VerifyJWT(token)
		if err != nil || !ok {
			log.WithError(err).Infof("verify jwt failed")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Set example variable
		ctx.Set("user", claims)

		// before request
		ctx.Next()
	}
}
