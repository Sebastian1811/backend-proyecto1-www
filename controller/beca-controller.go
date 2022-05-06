package controller

import (
	"strconv"

	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/Sebastian1811/backend-proyecto1-www/service"
	"github.com/gin-gonic/gin"
)

type BecaController interface {
	GetAll() []entity.Beca
	Save(ctx *gin.Context) entity.Beca
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.BecaService
}

func New(service service.BecaService) BecaController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll() []entity.Beca {
	return c.service.GetAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Beca {
	var beca entity.Beca
	ctx.BindJSON(&beca)
	c.service.Save(beca)
	return beca
}

func (c *controller) Update(ctx *gin.Context) error {
	var beca entity.Beca
	ctx.BindJSON(&beca)
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	beca.ID = id
	c.service.Update(beca)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var beca entity.Beca
	ctx.BindJSON(&beca)
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	beca.ID = id
	c.service.Delete(beca)
	return nil
}
