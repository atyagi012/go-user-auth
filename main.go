package main

import (
	"fmt"

	"github.com/atyagi012/go-user-auth/config"
	"github.com/atyagi012/go-user-auth/db"
	"github.com/atyagi012/go-user-auth/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Printf("Hello User")

	router := gin.Default()
	router.Use(gin.Logger())

	routes.UserRouter(router)

	fmt.Println("Server is running on port = ", config.Config.AppPort)
	router.Run(":" + config.Config.AppPort)
}

func init() {
	//fmt.Println("init db...")
	db.ConnectToDb()
	db.SyncDatabase()
}
