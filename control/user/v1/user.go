package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobackend/service"
	params "gobackend/structs/user"
	"gobackend/utils"
)

var userService service.UserService

func Login(c *gin.Context) {
	var login params.Login
	if err := c.ShouldBindJSON(&login); err == nil {
		res, err := userService.Login(login.Username, login.Password)
		if err == nil {
			utils.RespOk(c, res, "登陆成功")
			return
		} else {
			utils.RespFail(c, err.Error())
		}

	} else {
		utils.RespFail(c, err.Error())
	}
}

func Register(c *gin.Context) {
	var login params.Login
	if err := c.ShouldBindJSON(&login); err == nil {
		res, err := userService.Register(login.Username, login.Password)
		fmt.Println(res)
		if err == nil {
			utils.RespOk(c, res, "注册成功")
			return
		} else {
			utils.RespFail(c, err.Error())
		}

	} else {
		utils.RespFail(c, err.Error())
	}
}
