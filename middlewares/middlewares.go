package middlewares

import (
	"go-project/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthValidate() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := security.TokenValidate(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err})
			c.Abort()
			return
		}

		c.Next()
	}
}
