package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 注册
func UserRegister(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	// gin.key 取出服务实例
	userService := c.Keys["user"].(service.UserServiceClient)
	fmt.Println("userService:",userService)
	userResp, err := userService.UserRegister(context.Background(),&userReq)
	fmt.Println("userResp:",userResp,err)
	PanicIfUserError(err)
	
	r := res.Response{
		Status: uint(userResp.Code),
		Data:   userResp,
		Msg:    e.GetMsg(uint(userResp.Code)),
	}

	c.JSON(http.StatusOK,r)
}

// UserLogin 登录
func UserLogin(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))

	// gin.key 取出服务实例
	userService := c.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(),&userReq)
	PanicIfUserError(err)

	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Status: uint(userResp.Code),
		Data:   res.TokenData{
			User:  userResp.UserDetail,
			Token: token,
		},
		Msg:    e.GetMsg(uint(userResp.Code)),
	}

	c.JSON(http.StatusOK,r)
}
