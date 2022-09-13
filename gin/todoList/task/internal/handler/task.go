package handler

import (
	"context"
	"task/internal/repository"
	"task/internal/service"
	"task/pkg/e"
)

type TaskServer struct {
	
}

func NewTaskServer() *TaskServer {
	return &TaskServer{}
}

func (t *TaskServer) TaskCreate(ctx context.Context, req *service.TaskRequest) ( resp *service.CommonResponse, err error)  {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Success

	if err = task.TaskCreat(req); err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(e.Success)
	return

}

func (t *TaskServer)TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error)  {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Success

	if err = task.TaskUpdate(req);err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(e.Success)
	return
}

func (t *TaskServer) TaskShow(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse,err error)  {
	var task repository.Task
	resp = new(service.TasksDetailResponse)
	resp.Code = e.Success

	respT, err := task.TaskShow(req);
	if err != nil {
		resp.Code = e.Error
		resp.TaskDetail = nil
		return
	}

	resp.TaskDetail = repository.BuildTasks(respT)
	return resp, nil

}

func (t *TaskServer) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error)  {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Success

	if err = task.TaskDelete(req);err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		return
	}

	resp.Msg = e.GetMsg(e.Success)
	return resp, nil
}



