package service

import (
	"git.micro.com/changz/pruduct/domain/model"
	"git.micro.com/changz/pruduct/domain/repository"
)

type IProductRepository interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProductById(int64) error
	UpdateProduct(*model.Product) error
	FindProductById(int64) (*model.Product, error)
	FindAllProduct() ([]*model.Product, error)
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

// NewProductRepository 创建
func NewProductRepository(productRepository repository.IProductRepository) IProductRepository {
	return &ProductDataService{
		ProductRepository: productRepository,
	}
}

func (u *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return u.ProductRepository.CreateProduct(product)
}
func (u *ProductDataService) DeleteProductById(productID int64) error {
	return u.ProductRepository.DeleteProductByID(productID)
}
func (u *ProductDataService) UpdateProduct(product *model.Product) error {
	return u.ProductRepository.UpdateProduct(product)
}
func (u *ProductDataService) FindProductById(productID int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(productID)
}

func (u *ProductDataService) FindAllProduct() ([]*model.Product, error) {
	return u.ProductRepository.FindAll()
}
