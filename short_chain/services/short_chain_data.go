package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goproject/short_chain/global"
	"goproject/short_chain/models/table"
	"goproject/short_chain/msg"
	"goproject/short_chain/utils"
	"net/http"
)

type Service struct {
	table.ShortChainData
	Page int `json:"page"`
	Size int `json:"size"`
}

// Add 添加短链
func (s *Service) Add(c *gin.Context) {
	s.ShortId = utils.GetUuid()
	err := s.ShortChainData.Add(&s.ShortChainData)
	if err != nil {
		msg.Error(c, "添加短链失败")
	}
	msg.Success(c, "添加短链成功")
	return
}

// Update 修改短链
func (s *Service) Update(c *gin.Context) {
	shortQuery := table.ShortChainData{
		ShortId: s.ShortId,
	}

	err := shortQuery.QueryOne(&shortQuery)
	if err != nil {
		msg.Error(c, "编辑失败，短链不存在")
		return
	}

	shortQuery.MappingLinks = s.MappingLinks
	shortQuery.OriginalLink = s.OriginalLink
	db := global.Db
	err = db.Save(&shortQuery).Error
	if err != nil {
		msg.Error(c, "编辑短链失败")
		return
	}

	msg.Success(c, "编辑短链成功")
	return
}

// List 查询短链
func (s *Service) List(c *gin.Context) {
	//var list []table.ShortChainData
	var list *[]table.ShortChainData
	var count int64
	var err error

	db := global.Db
	db.Find(&list).Count(&count)
	err = db.Limit(s.Size).Offset((s.Page - 1) * s.Size).Order("created_at desc").Find(&list).Error
	if err != nil {
		msg.Error(c, "查询短链:"+err.Error())
		return
	}

	data := make([]map[string]interface{}, count)

	for k, v := range *list {
		item, _ := utils.ToMap(v, nil)
		data[k] = item
	}

	res := make(map[string]interface{})
	res["count"] = count
	res["list"] = data
	msg.Success(c, "success", res)
}

// Delete 删除短链
func (s *Service) Delete(c *gin.Context) {
	shortQuery := table.ShortChainData{
		ShortId: s.ShortId,
	}
	err := shortQuery.QueryOne(&shortQuery)
	if err != nil {
		msg.Error(c, "删除失败，短链不存在")
		return
	}

	db := global.Db
	err = db.Delete(&shortQuery).Error
	if err != nil {
		msg.Error(c, "删除短链失败")
		return
	}
	msg.Success(c, "删除成功")
	return
}

// Mapping 短链映射
func (s *Service) Mapping(c *gin.Context) {
	shortQuery := table.ShortChainData{
		OriginalLink: s.OriginalLink,
	}

	fmt.Println("shortQuery:", shortQuery)
	err := shortQuery.QueryOne(&shortQuery)
	if err != nil {
		msg.Error(c, "短链不存在")
		return
	}
	fmt.Println("Links:", shortQuery.MappingLinks)
	c.Redirect(http.StatusMovedPermanently, shortQuery.MappingLinks)
}
