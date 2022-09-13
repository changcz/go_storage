package msg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, msg string, data ...interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  msg,
			"data": data[0],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  msg,
			"data": data})
	}
	return
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  msg,
		"data": ""})
	return
}
