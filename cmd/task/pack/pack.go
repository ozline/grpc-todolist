package pack

import (
	"errors"

	"github.com/ozline/grpc-todolist/cmd/task/model"
	"github.com/ozline/grpc-todolist/pkg/errno"

	"github.com/ozline/grpc-todolist/idl/pb/task"
)

func BuildBaseResp(err error) *task.Base {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func BuildTask(source *model.Task) *task.Task {
	return &task.Task{
		Id:        source.ID,
		UserId:    source.UserId,
		Title:     source.Title,
		Status:    source.Status,
		Content:   source.Content,
		CreatedAt: source.CreatedAt.String(),
		UpdatedAt: source.UpdatedAt.String(),
	}
}

func BuildTaskList(tasks []*model.Task) []*task.Task {
	list := make([]*task.Task, 0, len(tasks))

	for _, info := range tasks {
		list = append(list, &task.Task{
			Id:        info.ID,
			UserId:    info.UserId,
			Title:     info.Title,
			Status:    info.Status,
			Content:   info.Content,
			CreatedAt: info.CreatedAt.String(),
			UpdatedAt: info.UpdatedAt.String(),
		})
	}

	return list
}

func baseResp(err errno.ErrNo) *task.Base {
	return &task.Base{
		Code:    err.ErrorCode,
		Message: err.ErrorMsg,
	}
}
