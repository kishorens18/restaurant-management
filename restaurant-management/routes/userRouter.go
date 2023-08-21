package routes

import (
	"restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUP())
	incomingRoutes.POST("/users/login", controllers.Login())
}
