package handler

import (
	"context"
	"git.micro.com/changz/cart/domain/model"
	"git.micro.com/changz/cart/domain/service"
	cart "git.micro.com/changz/cart/proto/cart"
	"git.micro.com/changz/common"
	log "github.com/micro/go-micro/v2/logger"
)

type Cart struct {
	CartDataService service.ICartRepository
}

// AddCart 添加购物车
func (c *Cart) AddCart(ctx context.Context, in *cart.CartInfo, out *cart.RespAddCart) error {
	cartInfo := &model.Cart{}
	err := common.SwapTo(in, cartInfo)
	if err != nil {
		log.Error("转换错误")
		return err
	}

	cartID, err := c.CartDataService.AddCart(cartInfo)
	if err != nil {
		log.Error("添加失败")
		return err
	}

	out.CartId = cartID
	out.Msg = "success"

	return nil
}

// CleanCart 清空购物车
func (c *Cart) CleanCart(ctx context.Context, in *cart.ReqClean, out *cart.Resp) error {
	err := c.CartDataService.CleanCart(in.UserId)
	if err != nil {
		log.Error("清空失败")
		return err
	}
	out.Msg = "success"

	return nil
}

// IncrCart 购物车 商品数量加
func (c *Cart) IncrCart(ctx context.Context, in *cart.ReqIncr, out *cart.Resp) error {
	err := c.CartDataService.IncrCart(in.CartId, in.ChangeNum)
	if err != nil {
		log.Error("商品数量增加失败")
		return err
	}
	out.Msg = "success"
	return nil
}

func (c *Cart) DecrCart(ctx context.Context, in *cart.ReqDecr, out *cart.Resp) error {
	err := c.CartDataService.DecrCart(in.CartId, in.ChangeNum)
	if err != nil {
		log.Error("商品数量减少失败")
		return err
	}
	out.Msg = "success"
	return nil
}

func (c *Cart) DeleteCartById(ctx context.Context, in *cart.ReqCartId, out *cart.Resp) error {
	err := c.CartDataService.DeleteCartById(in.CartId)
	if err != nil {
		log.Error("商品数量减少失败")
		return err
	}
	out.Msg = "success"
	return nil
}

// GetAll 获取购物车所有商品
func (c *Cart) GetAll(ctx context.Context, in *cart.ReqGetAll, out *cart.RespGetAll) error {
	allCart, err := c.CartDataService.FindAllCart(in.UserId)
	if err != nil {
		return err
	}
	cartToResponse(allCart, out)
	return nil
}

// cartToResponse 类型转换
func cartToResponse(cartSlice []*model.Cart, out *cart.RespGetAll) {
	for _, ps := range cartSlice {
		cartInfo := &cart.CartInfo{}
		err := common.SwapTo(ps, cartInfo)
		if err != nil {
			log.Error(err)
			break
		}
		out.CartInfo = append(out.CartInfo, cartInfo)
	}
}
