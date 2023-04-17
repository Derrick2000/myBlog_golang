package router

import (
	"myBlog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.RegisterUser)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	e.GET("/", controller.Index)

	e.GET("/blog_index", controller.GetBlogIndex)
	e.GET("/blog", controller.GoAddBlog)
	e.POST("/blog", controller.AddBlog)

	e.GET("/blog_detail", controller.BlogDetail)
	e.Run()
}
