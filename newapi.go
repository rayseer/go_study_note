package main

import (
	"./controller/user"
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func main() {
	//这里主要写路由
	fmt.Println("Start port 3000...")

	route := routes.New()
	//TODO: 有时候必须使用user.PostLogin()
	//有时候只能使用user.PostLogin
	//猜测：变量默认有值可省略，不然不可以省略
	route.Get("/:last/:first", user.PostLogin)

	http.Handle("/", route)
	http.ListenAndServe(":3000", nil)
}
