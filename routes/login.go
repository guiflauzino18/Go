package routes

import (
	"go-project/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRouters(r *gin.Engine) {

	r.POST("/login", controllers.Login)

}

