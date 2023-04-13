package controller

import (
	"myBlog/dao"
	"myBlog/model"

	"github.com/gin-gonic/gin"
)

func AddUser(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mngr.AddUser(&user)
}

func ListUser(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}
