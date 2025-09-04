package middlewares

import (
	"fmt"
	"go-project/security"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthValidate() gin.HandlerFunc {

	return func(c *gin.Context) {
		claims, err := security.TokenValidate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erro": err})
			c.Abort()
			return
		}

		//Recupera usuario e salva no contexto do gin
		username := claims["username"].(string)
		fmt.Println(claims)
		c.Set("username", username)
		c.Next()
	}
}

func ActionValidate(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("username")
		obj := c.FullPath()
		act := c.Request.Method

		fmt.Println(user)

		if !exists {
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
			c.JSON(http.StatusForbidden, gin.H{"erro": "Acesso Negado para " + fmt.Sprintf("%s", user)})
			c.Abort()
			return
		}

		c.Next()
	}
}
