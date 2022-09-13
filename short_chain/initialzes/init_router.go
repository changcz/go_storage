package initialzes

import (
	"github.com/gin-gonic/gin"
	"goproject/short_chain/router"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})
	router.ShortRouter(r)
	r.Run(":8001")
}
