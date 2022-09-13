package repository

import (
	"git.micro.com/changz/cart/domain/model"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/errors"
)

type ICartRepository interface {
	// InitTable 初始化数据表
	InitTable() error
	// AddCart 添加购物车
	AddCart(*model.Cart) (int64, error)
	// CleanCart 清空购物车
	CleanCart(int64) error
	// IncrCart 添加购物号商品数量
	IncrCart(int64, int64) error
	// DecrCart 减少购物号商品数量
	DecrCart(int64, int64) error
	// DeleteCartById 根据购物车id删除
	DeleteCartById(int64) error
	// GetAll 获取购物车所有商品
	GetAll(int64) ([]*model.Cart, error)

	// UpdateCart 更新购物车
	UpdateCart(*model.Cart) error
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

// NewCartRepository 创建
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}
func (c *CartRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) AddCart(cart *model.Cart) (cartID int64, err error) {
	// 先查询是否有记录
	db := c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}

	// 查看是否有受影响的行数
	if db.RowsAffected == 0 {
		// 建议使用自定义错误
		return 0, errors.New("10001", "购物车插入失败", 10001)
	}

	return cart.ID, nil
}

func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysqlDb.Where("id = ?", userID).Delete(&model.Cart{}).Error
}

func (c *CartRepository) DeleteCartById(cartID int64) error {
	return c.mysqlDb.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(&model.Cart{}).Update(cart).Error
}

func (c *CartRepository) GetAll(userID int64) (carts []*model.Cart, err error) {
	return carts, c.mysqlDb.Model(&model.Cart{}).Where("user_id = ?", userID).Find(&carts).Error
}

func (c *CartRepository) IncrCart(cartID, num int64) error {
	cart := &model.Cart{
		ID: cartID,
	}
	return c.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (c *CartRepository) DecrCart(cartID, num int64) error {
	cart := &model.Cart{
		ID: cartID,
	}

	// 判断数量是否够减
	db := c.mysqlDb.Model(cart).Where("num > ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("10002", "减少失败", 10002)
	}

	return nil
}
