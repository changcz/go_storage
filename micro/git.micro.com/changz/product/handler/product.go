package handler

import (
	"context"
	"git.micro.com/changz/pruduct/common"
	"git.micro.com/changz/pruduct/domain/model"
	"git.micro.com/changz/pruduct/domain/service"
	product "git.micro.com/changz/pruduct/proto/product" // 此处可以用别名可以用 .    （使用 in *ProductInfo ）
	"github.com/micro/go-micro/v2/util/log"
)

type Product struct {
	ProductDataService service.IProductRepository
}

// AddProduct 添加商品
func (p *Product) AddProduct(ctx context.Context, in *product.ProductInfo, out *product.ProductResponse) error {
	productInfo := &model.Product{}

	if err := common.SwapTo(in, productInfo); err != nil {
		return err
	}

	productID, err := p.ProductDataService.AddProduct(productInfo)
	if err != nil {
		return err
	}
	out.ProductId = productID
	return nil
}

// FindProductById 根据ID查找商品
func (p *Product) FindProductById(ctx context.Context, in *product.RequestProductId, out *product.ProductInfo) error {

	productInfo, err := p.ProductDataService.FindProductById(in.ProductId)
	if err != nil {
		return err
	}

	if err := common.SwapTo(productInfo, out); err != nil {
		return err
	}

	return nil
}

// UpdateProduct 修改商品
func (p *Product) UpdateProduct(ctx context.Context, in *product.ProductInfo, out *product.Response) error {
	productInfo := &model.Product{}

	if err := common.SwapTo(in, productInfo); err != nil {
		return err
	}

	if err := p.ProductDataService.UpdateProduct(productInfo); err != nil {
		return err
	}

	out.Msg = "success"
	return nil
}

// DeleteProductById 根据ID删除商品
func (p *Product) DeleteProductById(ctx context.Context, in *product.RequestProductId, out *product.Response) error {
	err := p.ProductDataService.DeleteProductById(in.ProductId)
	if err != nil {
		return err
	}
	out.Msg = "success"
	return nil
}

// FindAllProduct 获取所有商品
func (p *Product) FindAllProduct(ctx context.Context, in *product.RequestProductAll, out *product.ResponseProductAll) error {
	allProduct, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	productToResponse(allProduct, out)
	return nil
}

// productToResponse 类型转换
func productToResponse(productSlice []*model.Product, out *product.ResponseProductAll) {
	for _, ps := range productSlice {
		productInfo := &product.ProductInfo{}
		err := common.SwapTo(ps, productInfo)
		if err != nil {
			log.Error(err)
			break
		}
		out.ProductInfo = append(out.ProductInfo, productInfo)
	}
}
