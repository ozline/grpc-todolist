package rpc

import (
	"context"

	"github.com/ozline/grpc-todolist/idl/pb/task"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

func TaskCreate(ctx context.Context, req *task.CreateRequest) (*task.Task, error) {
	resp, err := TaskClient.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskUpdate(ctx context.Context, req *task.UpdateRequest) (*task.Task, error) {
	resp, err := TaskClient.Update(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func TaskDelete(ctx context.Context, req *task.DeleteRequest) error {
	resp, err := TaskClient.Delete(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return nil
}

func TaskGetList(ctx context.Context, req *task.GetListRequest) ([]*task.Task, error) {
	resp, err := TaskClient.GetList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}
