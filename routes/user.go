package routes

import (
	"go-project/controllers"
	"go-project/middlewares"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine, u controllers.UserController, e *casbin.Enforcer) {

	user := r.Group("/api/user", middlewares.AuthValidate(), middlewares.ActionValidate(e))
	//user := r.Group("/api/user")
	user.GET("/all", u.FindAll)
	user.GET("/:id", u.FindByID)
	user.POST("/create", u.Create)
}
