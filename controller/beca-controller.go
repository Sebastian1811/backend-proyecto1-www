package controller

import (
	"fmt"
	"strconv"

	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/service"
	"github.com/gin-gonic/gin"
)

type BecaController interface {
	GetAll(*gin.Context)
	Save(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetById(*gin.Context)
}

type controller struct {
	service service.BecaService
}

func New(service service.BecaService) BecaController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll(ctx *gin.Context) {
	ctx.JSON(200, c.service.GetAll())
}

func (c *controller) Save(ctx *gin.Context) {
	var beca entity.Beca
	ctx.BindJSON(&beca)
	c.service.Save(beca)
}

func (c *controller) Update(ctx *gin.Context) {
	var beca entity.Beca
	ctx.BindJSON(&beca)
	fmt.Println("hola")
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(404, err)
	} else {
		beca.ID = id
		c.service.Update(beca)
		ctx.JSON(200, beca)
	}
}

func (c *controller) Delete(ctx *gin.Context) {
	var beca entity.Beca
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithError(404, err)
	} else {
		beca = c.service.GetById(int(id))
		c.service.Delete(int(id))
		ctx.JSON(200, beca)
	}
}

func (c *controller) GetById(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 0)
	ctx.JSON(200, c.service.GetById(int(id)))
}
