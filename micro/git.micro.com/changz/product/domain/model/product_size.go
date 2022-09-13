package model

type ProductSize struct {
	ID            int64  `gorm:"primary_key;not_null;auto_increment" json:"id"` // 主键ID
	SizeName      string `json:"size_name"`                                     // 名称
	SizeCode      string `gorm:"unique_index:not_null" json:"size_code"`        // 标识
	SizeProductID int64  `json:"size_product_id"`
}
