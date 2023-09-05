package main

import (
	"github.com/gin-gonic/gin"

	"sample/controller"
	"sample/db"
	"sample/model/user"
	"sample/myauth"
)

func main() {
	db.GormConnect()
	defer db.GetDB().Close()

	// Migrate the schema
	db.GetDB().AutoMigrate(&user.User{})

	r := gin.Default()
	r.GET("/pong", controller.PingHandler)
	r.POST("/signup", controller.SignupHandler)
	users := r.Group("/users")
	users.Use(myauth.BasicAuthMiddleware())
	{
		users.GET("/:id", controller.GetUserHandler)
		users.PATCH("/:id", controller.PatchUserHandler)
	}
	close := r.Group("/close")
	close.Use(myauth.BasicAuthMiddleware())
	{
		close.POST("/", controller.CloseHandler)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
