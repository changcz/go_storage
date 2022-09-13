package model

type ProductSeo struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"` // 主键ID
	SeoTittle      string `json:"seo_tittle"`                                    // 标题
	SeoKeywords    string `json:"seo_keywords"`                                  // 关键字
	SeoDescription string `json:"seo_description"`                               // 描述
	SeoCode        string `gorm:"unique_index:not_null" json:"seo_code"`         // 标识
	SeoProductID   int64  `json:"seo_product_id"`
}
