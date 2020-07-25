package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)       // 创建用户
	router.POST("/user/:user_name", Login) // 登录
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

