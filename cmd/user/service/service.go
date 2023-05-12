package service

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/user/dal/db"
	"github.com/ozline/grpc-todolist/cmd/user/model"
	"github.com/ozline/grpc-todolist/idl/pb/user"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (us *UserService) Login(req *user.LoginRequest) (resp *model.User, err error) {
	err = db.CheckUserPassword(us.ctx, req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	return db.GetUserByUsername(us.ctx, req.Username)
}

func (us *UserService) Register(req *user.RegisterRequest) error {
	return db.CreateUser(us.ctx, req)
}
