package repository

import (
	"git.micro.service.category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// FindCategoryByID 根据分类id查找分类信息
	FindCategoryByID(int64) (*model.Category, error)
	// CreateCategory 创建分类
	CreateCategory(*model.Category) (int64, error)
	// DeleteCategoryByID 根据分类id删除分类
	DeleteCategoryByID(int64) error
	// UpdateCategory 更新分类信息
	UpdateCategory(*model.Category) error
	// FindAll 查找所有分类
	FindAll() ([]*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]*model.Category, error)
	FindCategoryByParent(int64) ([]*model.Category, error)
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// NewCategoryRepository 创建
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}
func (u *CategoryRepository) FindCategoryByID(categoryID int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.First(category, categoryID).Error
}
func (u *CategoryRepository) CreateCategory(category *model.Category) (categoryID int64, err error) {
	return category.ID, u.mysqlDb.Create(category).Error
}
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return u.mysqlDb.Where("id = ?", categoryID).Delete(&model.Category{}).Error
}
func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(&model.Category{}).Update(&category).Error
}
func (u *CategoryRepository) FindAll() (categorySlice []*model.Category, err error) {
	return categorySlice, u.mysqlDb.Model(&model.Category{}).Find(&categorySlice).Error
}

func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.Where("category_name = ?", categoryName).Find(&category).Error
}

func (u *CategoryRepository) FindCategoryByLevel(level uint32) (categorySlice []*model.Category, err error) {
	categorySlice = []*model.Category{}
	return categorySlice, u.mysqlDb.Where("category_level = ?", level).Find(&categorySlice).Error
}

func (u *CategoryRepository) FindCategoryByParent(parent int64) (categorySlice []*model.Category, err error) {
	categorySlice = []*model.Category{}
	return categorySlice, u.mysqlDb.Where("category_parent = ?", parent).Find(&categorySlice).Error
}
