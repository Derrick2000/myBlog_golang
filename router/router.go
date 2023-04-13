package router

import (
	"myBlog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/index", controller.ListUser)
	e.POST("/index", controller.AddUser)
	e.Run()
}
