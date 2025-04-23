package routes

import (
	"fmt"

	"github.com/atyagi012/go-user-auth/controllers"
	"github.com/atyagi012/go-user-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
	fmt.Println("I am in User router")
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.POST("/signup", controllers.CreateUser())
	incomingRoutes.POST("/login", controllers.Login())
	incomingRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)
}
