package routes

import (
	"go-project/controllers"
	"go-project/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine, u controllers.UserController) {

	user := r.Group("/user", middlewares.AuthValidate())
	user.GET("/all", u.FindAll)
	user.GET("/:id", u.FindByID)
	user.POST("/create", u.Create)
}
