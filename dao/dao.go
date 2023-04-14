package dao

import (
	"log"
	"myBlog/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	RegisterUser(user *model.User)
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
}

func (mngr *manager) RegisterUser(user *model.User) {
	mngr.db.Create(user)
}
