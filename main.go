package main

import (
	"API/db"
	"API/routers"
	"log"
)

func init() {
	//连接数据库
	db.InitDB()
}

func main() {
	router := routers.InitRouter()
	log.Fatal(router.Run())
}
