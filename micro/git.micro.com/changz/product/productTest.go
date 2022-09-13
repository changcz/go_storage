package main

import (
	"context"
	"fmt"
	"git.micro.com/changz/pruduct/common"
	git_micro_com_changz_product "git.micro.com/changz/pruduct/proto/product"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	//注册中心 //go get github.com/micro/go-plugins/registry/consul/v2
	consulR := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("git.micro.com.changz.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New service
	srv := micro.NewService(
		micro.Name("git.micro.com.changz.product.client"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8085"),
		// 添加consul 作为注册中心
		micro.Registry(consulR),

		// 绑定链路追踪 客户端是Client  服务端是 Handler
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := git_micro_com_changz_product.NewProductService("git.micro.com.changz.product", srv.Client())

	var productInfo = &git_micro_com_changz_product.ProductInfo{
		ProductName:        "mac book pro16",
		ProductSku:         "mac book pro16-2019",
		ProductPrice:       17850.1,
		ProductDescription: "苹果笔记本电脑",
		ProductCategoryId:  1,
		ProductImage: []*git_micro_com_changz_product.ProductImage{
			{
				ImageName: "mac-image-01",
				ImageCode: "mac-image-01",
				ImageUrl:  "http://127.0.0.1/image/01.jpg",
			},

			{
				ImageName: "mac-image-02",
				ImageCode: "mac-image-02",
				ImageUrl:  "http://127.0.0.1/image/02.jpg",
			},
		},
		ProductSize: []*git_micro_com_changz_product.ProductSize{
			{
				SizeName: "mac-size-01",
				SizeCode: "mac-size-01",
			},
			{
				SizeName: "mac-size-02",
				SizeCode: "mac-size-02",
			},
		},
		ProductSeo: &git_micro_com_changz_product.ProductSeo{
			SeoTittle:      "mac-book-pro tittle-01",
			SeoKeywords:    "mac-book-pro keywords-01",
			SeoDescription: "mac-book-pro description-01",
			SeoCode:        "mac-book-pro code-01",
		},
	}

	productID, err := productService.AddProduct(context.TODO(), productInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(productID)
}
