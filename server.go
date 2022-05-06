package main

import (
	"net/http"

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
	apiRoutes := server.Group("/Beca")
	{
		apiRoutes.POST("", func(ctx *gin.Context) {
			ctx.JSON(200, becaController.Save(ctx))
		})

		apiRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, becaController.GetAll())
		})

		apiRoutes.PUT(":id", func(ctx *gin.Context) {
			ctx.JSON(200, becaController.Update(ctx))
		})

		apiRoutes.DELETE(":id", func(ctx *gin.Context) {
			ctx.JSON(200, becaController.Delete(ctx))
		})

		apiRoutes.GET(":id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, becaController.GetById(ctx))
		})

	}

	server.Run("localhost:8080")
}
