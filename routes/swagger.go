package routes

import (
	"github.com/gin-gonic/gin"
)

func SwaggerRouters(r *gin.Engine) {

	r.GET("/swagger")

}
