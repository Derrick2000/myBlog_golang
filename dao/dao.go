package dao

import (
	"log"
	"myBlog/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	RegisterUser(user *model.User)
	Login(username string) model.User

	AddBlog(blog *model.Blog)
	GetAllBlogs() []model.Blog
	GetBlogById(pid int) model.Blog
}

type manager struct {
	db *gorm.DB
}

var Mngr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to init db", err)
	}
	Mngr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Blog{})
}

func (mngr *manager) RegisterUser(user *model.User) {
	mngr.db.Create(user)
}

func (mngr *manager) Login(username string) model.User {
	var user model.User
	mngr.db.Where("username = ?", username).First(&user)
	return user
}

func (mngr *manager) AddBlog(blog *model.Blog) {
	mngr.db.Create(blog)
}

func (mngr *manager) GetAllBlogs() []model.Blog {
	var blogs = make([]model.Blog, 10)
	mngr.db.Find(&blogs)
	return blogs
}

func (mngr *manager) GetBlogById(pid int) model.Blog {
	var blog model.Blog
	mngr.db.First(&blog, pid)
	return blog
}
