package model

type Product struct {
	ID                 int64          `gorm:"primary_key;not_null;auto_increment" json:"id"`  // 主键
	ProductName        string         `json:"product_name"`                                   // 分类名称
	ProductSku         string         `gorm:"unique_index:not_null" json:"product_sku"`       // 分类等级
	ProductPrice       float64        `json:"product_price"`                                  // 分类父类
	ProductDescription string         `json:"product_description"`                            // 分类描述
	ProductSeo         ProductSeo     `gorm:"ForeignKey:SeoProductID" json:"product_seo"`     // 分类照片
	ProductImage       []ProductImage `gorm:"ForeignKey:ImageProductID" json:"product_image"` // 分类照片
	ProductSize        []ProductSize  `gorm:"ForeignKey:SizeProductID" json:"product_size"`   // 分类照片
}
