package model

type ProductImage struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"` // 主键
	ImageName      string `json:"image_name"`                                    // 图片名称
	ImageCode      string `gorm:"unique_index:not_null" json:"image_code"`       // 图片标识
	ImageUrl       string `json:"image_url"`                                     // 图片url
	ImageProductID int64  `json:"image_product_id"`
}
