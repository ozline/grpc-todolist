package db

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/task/model"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/idl/pb/task"
)

func Create(ctx context.Context, req *task.CreateRequest) (*model.Task, error) {

	info := &model.Task{
		ID:      SF.NextVal(),
		UserId:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}

	err := DB.Create(info).Error

	if err != nil {
		return nil, err
	}

	return info, nil
}

func Delete(ctx context.Context, req *task.DeleteRequest) error {

	err := DB.Table(config.Service.Name).Where("id = ?", req.Id).Unscoped().Delete(&model.Task{}).Error

	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, req *task.UpdateRequest) (*model.Task, error) {
	var info *model.Task
	err := DB.Table(config.Service.Name).Where("id = ?", req.Id).First(&info).Error

	if err != nil {
		return nil, err
	}

	info.Content = req.Content
	info.Title = req.Title
	info.Status = req.Status

	err = DB.Table(config.Service.Name).Save(info).Error

	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetList(ctx context.Context, req *task.GetListRequest) ([]*model.Task, error) {
	var list []*model.Task

	err := DB.Table(config.Service.Name).Where("user_id = ? and status = ?", req.UserId, req.Status).Find(&list).Error

	if err != nil {
		return nil, err
	}

	return list, nil
}
