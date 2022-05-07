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
	apiRoutes := server.Group("/Beca")
	{
		apiRoutes.POST("", becaController.Save)

		apiRoutes.GET("/all", becaController.GetAll)

		apiRoutes.PUT(":id", becaController.Update)

		apiRoutes.DELETE(":id", becaController.Delete)

		apiRoutes.GET(":id", becaController.GetById)
	}
	server.Run("localhost:8080")
}
