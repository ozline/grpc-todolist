package rpc

import (
	"context"

	"github.com/ozline/grpc-todolist/idl/pb/user"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

func UserLogin(ctx context.Context, req *user.LoginRequest) (*user.User, error) {
	resp, err := UserClient.Login(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return resp.Data, nil
}

func UserRegister(ctx context.Context, req *user.RegisterRequest) error {
	resp, err := UserClient.Register(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Message)
	}

	return nil
}
