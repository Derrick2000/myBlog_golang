package main

import (
	"myBlog/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router.Start()
}
