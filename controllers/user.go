package controllers

import (
	"go-project/model"
	"go-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (u *UserController) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": err})
		return
	}

	err := user.Prepare("create")

	id, err := u.service.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userid": id})
}

func (u *UserController) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": err})
		return
	}

	user, err := u.service.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) FindAll(c *gin.Context) {
	users, err := u.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	c.JSON(http.StatusOK, users)

}
