package common

import "github.com/gin-gonic/gin"

// 200,请求成功
func Ok(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

// 自定义错误
func Err(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"code": -1,
		"msg":  err.Error(),
	})
}

// 没有匹配路由
func NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{
		"code": -1,
		"msg":  "No route found",
	})
}

// 400,请求报文存在语法错误
func Err400(c *gin.Context, err error) {
	c.JSON(400, err.Error())
}

// 500,表示服务器端在执行请求时发生了错误
func Err500(c *gin.Context, err error) {
	c.JSON(500, err.Error())
}
