package router

import (
	"github.com/gin-gonic/gin"
	user_v1 "gobackend/control/user/v1"
)

func Init(e *gin.Engine) {
	user := e.Group("/user")
	{
		v1 := user.Group("/v1")
		v1.POST("/login", user_v1.Login)
		v1.POST("/register", user_v1.Register)
		v1.GET("/info", user_v1.UserInfo)
		v1.POST("/logout", user_v1.LoginOut)
	}
}
