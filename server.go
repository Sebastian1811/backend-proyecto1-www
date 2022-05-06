package main

import (
	"github.com/Sebastian1811/backend-proyecto1-www/controller"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
	"github.com/Sebastian1811/backend-proyecto1-www/service"
	"github.com/gin-gonic/gin"
)

var (
	becaRepository repository.BecaRepository = repository.NewBecaRepository()
	becaService    service.BecaService       = service.New(becaRepository)
	becaController controller.BecaController = controller.New(becaService)
)

func main() {
	server := gin.Default()
	server.POST("/Beca", func(ctx *gin.Context) {
		ctx.JSON(200, becaController.Save(ctx))
	})

	server.GET("/Becas", func(ctx *gin.Context) {
		ctx.JSON(200, becaController.GetAll())
	})

	server.PUT("/Beca/:id", func(ctx *gin.Context) {
		ctx.JSON(200, becaController.Update(ctx))
	})

	server.DELETE("/Beca/:id", func(ctx *gin.Context) {
		ctx.JSON(200, becaController.Delete(ctx))
	})

	server.Run("localhost:8080")
}
