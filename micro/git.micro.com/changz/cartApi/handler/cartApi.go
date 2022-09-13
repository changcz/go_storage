package handler

import (
	"context"
	"errors"
	"fmt"

	log "github.com/micro/go-micro/v2/logger"
	"strconv"

	api "github.com/micro/go-micro/v2/api/proto"
)

// CartApi CartService
type CartApi struct{

}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// CartApi.FindAll is called by the API as /cartApi/FindAll with post body {"name": "foo"}\

func (e *CartApi) FindAll(ctx context.Context, req *api.Request, resp *api.Response) error {
	log.Info("接收到/cartApi/FindAll访问请求参数")

	if _,ok  := req.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}

	userIdStr := req.Get["user_id"].Values[0]
	fmt.Println(userIdStr)

	userID, err := strconv.ParseInt(userIdStr, 64, 10)
	if err != nil {
		return err
	}

	// 根据userID获取购物车商品
	//e.cartService.GetAll(context.TODO())
	resp.StatusCode = 200
	resp.Body = strconv.FormatInt(userID, 10)
	return nil
}