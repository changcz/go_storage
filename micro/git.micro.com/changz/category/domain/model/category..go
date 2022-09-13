package model

type Category struct {
	ID                  int64  `gorm:"primary_key;not_null;auto_increment" json:"id"` // 主键
	CategoryName        string `gorm:"unique_index.not_null" json:"category_name"`    // 分类名称
	CategoryLevel       uint32 `json:"category_level"`                                // 分类等级
	CategoryParent      int64  `json:"category_parent"`                               // 分类父类
	CategoryImage       string `json:"category_image"`                                // 分类照片
	CategoryDescription string `json:"category_description"`                          // 分类描述
}
