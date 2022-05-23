package controller

import (
	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user entity.User

	password := []byte(ctx.Param("password"))

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	ctx.BindJSON(&user)
	user.Password = string(hashedPassword)
	c.service.Register(user)
}
func (c *userController) Login(ctx *gin.Context) {

}
