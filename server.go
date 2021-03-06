package main

import (
	"github.com/Sebastian1811/backend-proyecto1-www/config"
	"github.com/Sebastian1811/backend-proyecto1-www/controller"
	"github.com/Sebastian1811/backend-proyecto1-www/middlewares"
	"github.com/Sebastian1811/backend-proyecto1-www/repository"
	"github.com/Sebastian1811/backend-proyecto1-www/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConn()
	requisitoRepository repository.RequisitoRepository = repository.NewRequesitoRepository(db)
	requisitoService    service.RequisitoService       = service.NewRequisitoService(requisitoRepository)

	becaRepository repository.BecaRepository = repository.NewBecaRepository(db)
	becaService    service.BecaService       = service.New(becaRepository)
	becaController controller.BecaController = controller.New(becaService, requisitoService)

	authService    service.JWTService        = service.NewJWTService()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService, authService)
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())

	apiRoutes := server.Group("/Beca", middlewares.AuthorizeJWT())
	{
		apiRoutes.POST("", becaController.Save)

		apiRoutes.GET("/all", becaController.GetAll)

		apiRoutes.PUT(":id", becaController.Update)

		apiRoutes.DELETE(":id", becaController.Delete)

		apiRoutes.GET(":id", becaController.GetById)

		apiRoutes.GET("/categoria/:categoria", becaController.GetByCategoria)
	}

	userRoutes := server.Group("/User")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}

	server.Run(":8080")
}
