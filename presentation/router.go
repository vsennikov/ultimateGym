package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/infrastructure"
	"github.com/vsennikov/ultimateGym/presentation/controllers"
	"github.com/vsennikov/ultimateGym/services"
)


type Router struct {
	UserController *controllers.UserController
}

func NewRouter(u *controllers.UserController) *Router {
	return &Router{
		UserController: u,
	}
}

func InitController() {
	userDB := infrastructure.UserDB{}
	userService := services.NewUserService(&userDB)
	userController := controllers.NewUserController(userService)
	
	router := NewRouter(userController)
	router.InitRouter()
}

func (r *Router) InitRouter() {
	router := gin.Default()

	api :=router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/registration", r.UserController.Registration)
		}
		tg := api.Group("/tg")
		{
			tg.POST("/registration", r.UserController.TelegramRegistration)
		}
	}
	router.Run()
}
