package main

import (
	"./controller/user"
	"./route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start port 3000...")

	//用户登陆
	//curl "localhost:3000/user/login"
	route.Post("/user/login", user.PostLogin)
	//获取用户简介
	//curl "localhost:3000/user/profile"
	route.Get("/user/profile", user.GetProfile())

	log.Fatal(http.ListenAndServe(":3000", nil))
}
