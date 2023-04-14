package controller

import (
	"myBlog/dao"
	"myBlog/model"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mngr.RegisterUser(&user)
	// context.Redirect(200, "/")
}

func GoRegister(context *gin.Context) {
	context.HTML(200, "register.html", nil)
}

func Index(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}

func ListUser(context *gin.Context) {
	context.HTML(200, "userlist.html", nil)
}
