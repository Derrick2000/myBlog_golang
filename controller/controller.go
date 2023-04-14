package controller

import (
	"fmt"
	"myBlog/dao"
	"myBlog/model"

	"github.com/gin-gonic/gin"
)

func GoRegister(context *gin.Context) {
	context.HTML(200, "register.html", nil)
}

func RegisterUser(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mngr.RegisterUser(&user)
	context.Redirect(301, "/")
}

func GoLogin(context *gin.Context) {
	context.HTML(200, "login.html", nil)
}

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	fmt.Println(username)
	u := dao.Mngr.Login(username)

	if u.Username == "" {
		context.HTML(200, "login.html", "user not found")
	} else {
		if u.Password != password {
			fmt.Println("Incorrect password")
			context.HTML(200, "login.html", "Incorrect password")
		} else {
			fmt.Println("Login Success")
			context.Redirect(301, "/")
		}
	}
}

func Index(context *gin.Context) {
	context.HTML(200, "index.html", nil)
}

func ListUser(context *gin.Context) {
	context.HTML(200, "userlist.html", nil)
}
