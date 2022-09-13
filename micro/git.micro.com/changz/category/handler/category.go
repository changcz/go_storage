package handler

import (
	"context"
	"git.micro.service.category/common"
	"git.micro.service.category/domain/model"
	"git.micro.service.category/domain/service"
	category "git.micro.service.category/proto/category"
	"github.com/micro/go-micro/v2/util/log"
)

type Category struct {
	CategoryDataService service.ICategoryRepository
}

// CreateCategory 创建分类
func (c *Category) CreateCategory(ctx context.Context, in *category.CategoryRequest, out *category.CreateCategoryResponse) error {
	categoryU := &model.Category{}

	err := common.SwapTo(in, categoryU)
	if err != nil {
		return err
	}

	id, err := c.CategoryDataService.AddCategory(categoryU)
	if err != nil {
		return err
	}
	out.Message = "success"
	out.CategoryId = id
	return nil
}

// UpdateCategory 分类更新
func (c *Category) UpdateCategory(ctx context.Context, in *category.CategoryRequest, out *category.UpdateCategoryResponse) error {
	categoryU := &model.Category{}
	err := common.SwapTo(in, categoryU)
	if err != nil {
		return err
	}

	err = c.CategoryDataService.UpdateCategory(categoryU)
	if err != nil {
		return err
	}
	out.Message = "success"
	return nil
}

// DeleteCategory 删除分类
func (c *Category) DeleteCategory(ctx context.Context, in *category.DeleteCategoryRequest, out *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(in.CategoryId)
	if err != nil {
		return err
	}
	out.Message = "success"
	return nil
}

// FindCategoryByName 根据分类名查找分类
func (c *Category) FindCategoryByName(ctx context.Context, in *category.FindByNameRequest, out *category.CategoryResponse) error {
	categoryU, err := c.CategoryDataService.FindCategoryByName(in.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryU, out)
}

// FindCategoryById 根据分类ID查找分类
func (c *Category) FindCategoryById(ctx context.Context, in *category.FindByIdRequest, out *category.CategoryResponse) error {
	categoryU, err := c.CategoryDataService.FindCategoryByID(in.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryU, out)
}

// FindCategoryByLevel 根据分类等级查找分类
func (c *Category) FindCategoryByLevel(ctx context.Context, in *category.FindByLevelRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(in.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil

}

// FindCategoryByParent 根据分类父类查找分类
func (c *Category) FindCategoryByParent(ctx context.Context, in *category.FindByParentRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(in.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil
}

// FindAllCategory 获取所有分类
func (c *Category) FindAllCategory(ctx context.Context, in *category.FindAllRequest, out *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, out)
	return nil
}

// categoryToResponse 类型转换
func categoryToResponse(categorySlice []*model.Category, out *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}

		out.Category = append(out.Category, cr)
	}
}
