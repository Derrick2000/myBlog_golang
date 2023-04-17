package controller

import (
	"fmt"
	"html/template"
	"myBlog/dao"
	"myBlog/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
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

func GetBlogIndex(context *gin.Context) {
	blogs := dao.Mngr.GetAllBlogs()
	context.HTML(200, "blogIndex.html", blogs)
}

func AddBlog(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	tag := context.PostForm("tag")

	blog := model.Blog{
		Title:   title,
		Content: content,
		Tag:     tag,
	}
	dao.Mngr.AddBlog(&blog)

	context.Redirect(302, "/blog_index")
}

func GoAddBlog(context *gin.Context) {
	context.HTML(200, "blog.html", nil)
}

func BlogDetail(context *gin.Context) {
	s := context.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mngr.GetBlogById(pid)

	content := blackfriday.Run([]byte(p.Content), blackfriday.WithNoExtensions())

	fmt.Printf("%s", template.HTML(content))
	context.HTML(200, "blogDetail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
