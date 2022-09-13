package router

import (
	"github.com/gin-gonic/gin"
	"goproject/short_chain/controllers"
)

func ShortRouter(r *gin.Engine) {
	r.POST("/add", controllers.Add)
	r.POST("/list", controllers.List)
	r.POST("/update", controllers.Update)
	r.POST("/delete", controllers.Delete)
	r.GET("/short/:original_link", controllers.Mapping)
}
