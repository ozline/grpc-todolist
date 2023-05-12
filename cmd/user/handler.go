package main

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/user/pack"
	"github.com/ozline/grpc-todolist/cmd/user/service"
	"github.com/ozline/grpc-todolist/idl/pb/user"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

type UserServiceImpl struct {
	user.UnimplementedUserServiceServer
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (*UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	info, err := service.NewUserService(ctx).Login(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = pack.BuildUser(info)
	return resp, nil
}

func (*UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	err = service.NewUserService(ctx).Register(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
