package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/ozline/grpc-todolist/cmd/user/dal"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/idl/pb/user"
)

func TestUserService_Login(t *testing.T) {
	config.Init("../../../config", "user", 0)
	dal.Init()
	ctx := context.Background()
	userService := NewUserService(ctx)
	r, err := userService.Login(&user.LoginRequest{
		Username: "FanOne",
		Password: "FanOne404",
	})
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("r", r)
}
