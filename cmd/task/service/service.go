package service

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/task/dal/db"
	"github.com/ozline/grpc-todolist/cmd/task/model"
	"github.com/ozline/grpc-todolist/idl/pb/task"
)

type TaskService struct {
	ctx context.Context
}

func NewTaskService(ctx context.Context) *TaskService {
	return &TaskService{ctx: ctx}
}

func (us *TaskService) Create(req *task.CreateRequest) (*model.Task, error) {
	return db.Create(us.ctx, req)
}

func (us *TaskService) Update(req *task.UpdateRequest) (*model.Task, error) {
	return db.Update(us.ctx, req)
}

func (us *TaskService) Delete(req *task.DeleteRequest) error {
	return db.Delete(us.ctx, req)
}
func (us *TaskService) GetList(req *task.GetListRequest) ([]*model.Task, error) {
	return db.GetList(us.ctx, req)
}
