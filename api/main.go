package main

import "github.com/gin-gonic/gin"

func RegisterHandlers() *gin.Engine {
	r := gin.Default()

	//users handeler
	r.POST("/user", RegisterUser)
	r.POST("/user/:user_name", Login)
	r.GET("/user/:user_name", GetUserInfo)
	r.DELETE("/user/:user_name", Logout)
	r.PUT("/user/:user_name/pwd/modify", ModifyPwd)
	r.PUT("/user/:user_name", ModifyUserInfo)

	return r
}

func main() {
	r := RegisterHandlers()

	r.Run(":8000")
}