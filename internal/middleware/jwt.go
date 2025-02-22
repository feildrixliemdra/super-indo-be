package middleware

import (
	"net/http"
	"super-indo-be/internal/util"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.ValidateJWT(c, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
