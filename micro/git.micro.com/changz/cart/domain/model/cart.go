package model

type Cart struct {
	ID        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"` // 主键
	ProductID int64 `gorm:"not_null" json:"product_id"`                    // 商品id
	Num       int64 `gorm:"not_null" json:"num"`                           // 数量
	SizeID    int64 `gorm:"not_null" json:"size_id"`                       // 类型id
	UserID    int64 `gorm:"not_null" json:"user_id"`                       // 用户id
}
