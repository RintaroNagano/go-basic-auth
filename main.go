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
	r.GET("/ping", controller.PingHandler)
	r.GET("/new", controller.NewHandler)
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

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
