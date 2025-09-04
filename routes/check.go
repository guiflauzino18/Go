package routes

import (
	"go-project/controllers"

	"github.com/gin-gonic/gin"
)

func CheckRouters(r *gin.Engine) {

	r.POST("/login", controllers.Check)

}
