package repository

import (
	"git.micro.com/changz/pruduct/domain/model"
	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindProductByID 根据分类id查找分类信息
	FindProductByID(int64) (*model.Product, error)
	// CreateProduct 创建分类
	CreateProduct(*model.Product) (int64, error)
	// DeleteProductByID 根据分类id删除分类
	DeleteProductByID(int64) error
	// UpdateProduct 更新分类信息
	UpdateProduct(*model.Product) error
	// FindAll 查找所有分类
	FindAll() ([]*model.Product, error)
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

// NewProductRepository 创建
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSeo{}, &model.ProductSize{}).Error
}
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, productID).Error
}
func (u *ProductRepository) CreateProduct(product *model.Product) (productID int64, err error) {
	return product.ID, u.mysqlDb.Create(product).Error
}
func (u *ProductRepository) DeleteProductByID(productID int64) (err error) {
	// 开启事务
	tx := u.mysqlDb.Begin()

	// 捕捉错误 回滚
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 删除 product 表
	err = tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除 image 表
	err = tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除 size 表
	err = tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除 seo 表
	err = tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDb.Model(&model.Product{}).Update(&product).Error
}
func (u *ProductRepository) FindAll() (productSlice []*model.Product, err error) {
	return productSlice, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Model(&model.Product{}).Find(&productSlice).Error
}
