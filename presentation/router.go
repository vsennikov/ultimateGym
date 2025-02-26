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
	ExerciseController *controllers.ExerciseController
}

func NewRouter(u *controllers.UserController, l *controllers.LoginController, e *controllers.ExerciseController) *Router {
	return &Router{
		UserController: u,
		LoginController: l,
		ExerciseController: e,
	}
}

func InitController() {
	userDB := infrastructure.UserDB{}
	exercisesDB := infrastructure.ExerciseDB{}

	userService := services.NewUserService(&userDB)
	exercisesService := services.NewExerciseService(&exercisesDB)

	userController := controllers.NewUserController(userService)
	loginController := controllers.NewLoginController(userService)
	exerciseController := controllers.NewExerciseController(exercisesService, userService)

	router := NewRouter(userController, loginController, exerciseController)
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
			v1.POST("/exercise", r.ExerciseController.CreateExercise)
			v1.GET("/exercises", r.ExerciseController.GetAllExercises)
			v1.GET("/user/exercises", r.ExerciseController.GetAllUserExercises)
			v1.GET("/exercises/:muscle_group", r.ExerciseController.GetAllExercisesByType)
			v1.GET("/user/exercises/:muscle_group", r.ExerciseController.GetAllUserExercisesByType)
			v1.GET("/exercise/:name", r.ExerciseController.GetExerciseByName)
			v1.DELETE("/exercise/:name", r.ExerciseController.DeleteExercise)
			v1.PUT("/exercise/:name", r.ExerciseController.UpdatedExercise)
		}
	}
	router.Run()
}
