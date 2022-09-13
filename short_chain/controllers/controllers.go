package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goproject/short_chain/msg"
	"goproject/short_chain/services"
)

// Add 添加短链
func Add(c *gin.Context) {
	var dataService services.Service
	err := c.ShouldBindJSON(&dataService)
	if err != nil {
		msg.Error(c, "请求参数错误:"+err.Error())
		return
	}

	if dataService.MappingLinks == "" {
		msg.Error(c, "原始链接不能为空")
		return
	}

	if dataService.OriginalLink == "" {
		msg.Error(c, "隐射链接不能为空")
		return
	}

	dataService.Add(c)
}

// Update 添加短链
func Update(c *gin.Context) {
	var dataService services.Service
	err := c.ShouldBindJSON(&dataService)
	if err != nil {
		msg.Error(c, "请求参数错误:"+err.Error())
		return
	}

	if dataService.MappingLinks == "" {
		msg.Error(c, "原始链接不能为空")
		return
	}

	if dataService.OriginalLink == "" {
		msg.Error(c, "隐射链接不能为空")
		return
	}

	if dataService.ShortId == "" {
		msg.Error(c, "短链id不能为空")
		return
	}
	dataService.Update(c)
	return
}

// List 查询短链列表
func List(c *gin.Context) {
	var dataService services.Service
	err := c.ShouldBindJSON(&dataService)
	if err != nil {
		msg.Error(c, "请求参数错误:"+err.Error())
		return
	}
	if dataService.Page == 0 {
		dataService.Page = 1
	}
	if dataService.Size == 0 {
		dataService.Size = 10
	}
	dataService.List(c)
	return
}

// Delete 删除短链
func Delete(c *gin.Context) {
	var dataService services.Service
	err := c.ShouldBindJSON(&dataService)
	if err != nil {
		msg.Error(c, "请求参数错误:"+err.Error())
		return
	}

	if dataService.ShortId == "" {
		msg.Error(c, "短链id不能为空")
		return
	}
	dataService.Delete(c)
	return
}

// Mapping 映射短链
func Mapping(c *gin.Context) {
	originalLink := c.Param("original_link")
	fmt.Println("ctl originalLink:", originalLink)
	if originalLink == "" {
		msg.Error(c, "参数不能为空")
		return
	}

	var dataService services.Service
	dataService.OriginalLink = originalLink
	dataService.Mapping(c)
}
