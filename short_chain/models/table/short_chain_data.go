package table

import (
	"goproject/short_chain/models"
	"gorm.io/gorm"
)

type ShortChainData struct {
	models.Crud
	gorm.Model
	ShortId      string `gorm:"size:255;comment:'短链id'" json:"short_id"`
	OriginalLink string `gorm:"text;comment:'原始链接'" json:"original_link"`
	MappingLinks string `gorm:"size:255;comment:'映射链接'" json:"mapping_links"`
}

func (ShortChainData) TableName() string {
	return "short_chain_data"
}
