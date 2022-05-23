package controller

import (
	"net/http"

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
	service    service.UserService
	jwtService service.JWTService
}

func NewUserController(service service.UserService, authService service.JWTService) UserController {
	return &userController{
		service:    service,
		jwtService: authService,
	}
}

func (c *userController) Register(ctx *gin.Context) {

	var user entity.User
	ctx.BindJSON(&user)
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	c.service.Register(user)
}
func (c *userController) Login(ctx *gin.Context) {
	var user entity.User
	ctx.BindJSON(&user)
	givenPass := user.Password
	email := user.Email
	user = c.service.Login(email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(givenPass))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "usuario o password incorrecta")

	} else {
		token := c.jwtService.GenerateToken()
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
