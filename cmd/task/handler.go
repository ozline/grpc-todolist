package main

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/task/pack"
	"github.com/ozline/grpc-todolist/cmd/task/service"
	"github.com/ozline/grpc-todolist/idl/pb/task"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

type TaskServiceImpl struct {
	task.UnimplementedTaskServiceServer
}

func NewTaskServiceImpl() *TaskServiceImpl {
	return &TaskServiceImpl{}
}

func (*TaskServiceImpl) Create(ctx context.Context, req *task.CreateRequest) (resp *task.CreateResponse, err error) {
	resp = new(task.CreateResponse)

	info, err := service.NewTaskService(ctx).Create(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}

func (*TaskServiceImpl) Update(ctx context.Context, req *task.UpdateRequest) (resp *task.UpdateResponse, err error) {
	resp = new(task.UpdateResponse)

	info, err := service.NewTaskService(ctx).Update(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTask(info)
	return resp, nil
}

func (*TaskServiceImpl) Delete(ctx context.Context, req *task.DeleteRequest) (resp *task.DeleteResponse, err error) {
	resp = new(task.DeleteResponse)

	err = service.NewTaskService(ctx).Delete(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

func (*TaskServiceImpl) GetList(ctx context.Context, req *task.GetListRequest) (resp *task.GetListResponse, err error) {
	resp = new(task.GetListResponse)

	list, err := service.NewTaskService(ctx).GetList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildTaskList(list)
	return resp, nil
}
