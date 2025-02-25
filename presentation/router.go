package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/infrastructure"
	"github.com/vsennikov/ultimateGym/presentation/controllers"
	"github.com/vsennikov/ultimateGym/services"
)


type Router struct {
	UserController *controllers.UserController
	LoginController *controllers.LoginController
}

func NewRouter(u *controllers.UserController, l *controllers.LoginController) *Router {
	return &Router{
		UserController: u,
		LoginController: l,
	}
}

func InitController() {
	userDB := infrastructure.UserDB{}
	userService := services.NewUserService(&userDB)
	userController := controllers.NewUserController(userService)
	loginController := controllers.NewLoginController(userService)
	
	router := NewRouter(userController, loginController)
	router.InitRouter()
}

func (r *Router) InitRouter() {
	router := gin.Default()

	api :=router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/registration", r.UserController.Registration)
			v1.POST("/login", r.LoginController.Login)
		}
		tg := api.Group("/tg")
		{
			tg.POST("/registration", r.UserController.TelegramRegistration)
		}
	}
	router.Run()
}
