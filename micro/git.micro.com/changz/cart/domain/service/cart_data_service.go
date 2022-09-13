package service

import (
	"git.micro.com/changz/cart/domain/model"
	"git.micro.com/changz/cart/domain/repository"
)

type ICartRepository interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCartById(int64) error
	UpdateCart(*model.Cart) error
	FindAllCart(int64) ([]*model.Cart, error)
	// IncrCart 添加购物号商品数量
	IncrCart(int64, int64) error
	// DecrCart 减少购物号商品数量
	DecrCart(int64, int64) error
	// CleanCart 清空购物车
	CleanCart(int64) error
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

// NewCartRepository 创建
func NewCartRepository(cartRepository repository.ICartRepository) ICartRepository {
	return &CartDataService{
		CartRepository: cartRepository,
	}
}

func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.AddCart(cart)
}
func (u *CartDataService) DeleteCartById(cartID int64) error {
	return u.CartRepository.DeleteCartById(cartID)
}
func (u *CartDataService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

func (u *CartDataService) FindAllCart(userID int64) ([]*model.Cart, error) {
	return u.CartRepository.GetAll(userID)
}

func (u *CartDataService) IncrCart(cartID, num int64) error {
	return u.CartRepository.IncrCart(cartID, num)
}

func (u *CartDataService) DecrCart(cartID, num int64) error {
	return u.CartRepository.IncrCart(cartID, num)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}
