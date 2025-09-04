package middlewares

import (
	"go-project/security"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthValidate() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := security.TokenValidate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ActionVAlidate(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, existis := c.Get("username")
		obj := c.FullPath()
		act := c.Request.Method

		if !existis {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": "Usuário não autenticado"})
			c.Abort()
			return
		}

		authorized, err := e.Enforce(user, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro Casbin"})
			c.Abort()
			return
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"erro": "Acesso Negado"})
			c.Abort()
			return
		}

		c.Next()
	}
}
