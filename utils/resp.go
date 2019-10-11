package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespFail(c *gin.Context, msg string) {
	Resp(c, -1, nil, msg)
}

func RespOk(c *gin.Context, data interface{}, msg string) {
	Resp(c, 0, data, msg)
}

func RespOkList(c *gin.Context, lists interface{}, total interface{}) {
	//分页数目,
	RespList(c, 0, lists, total)
}

func Resp(c *gin.Context, code int, data interface{}, msg string) {
	//设置header为JSON（默认是text/html）
	c.Header("Content-Type", "appliction/json")
	c.Header("Access-Control-Allow-Origin", "*")             //允许访问所有域
	c.Header("Access-Control-Allow-Headers", "Content-Type") //header的类型
	c.AbortWithStatus(http.StatusOK)

	//输出
	c.JSON(200, gin.H{
		"Code": code,
		"Msg":  msg,
		"Data": data,
	})
}

func RespList(c *gin.Context, code int, data interface{}, total interface{}) {

	c.Header("Content-Type", "appliction/json")
	//设置200状态
	c.AbortWithStatus(http.StatusOK)
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20

	//输出
	c.JSON(200, gin.H{
		"Code":  code,
		"Rows":  data,
		"Total": total,
	})
}
