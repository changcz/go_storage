package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/util"
)

func GetTaskList(ginCtx *gin.Context) {

	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskCreate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   taskResp,
		Status: uint(taskResp.Code),
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserID)
	TaskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	TaskResp, err := TaskService.TaskDelete(context.Background(),  &tReq)
	PanicIfTaskError(err)
	r := res.Response{
		Data:   TaskResp,
		Status: uint(TaskResp.Code),
		Msg:    e.GetMsg(uint(TaskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}