package db

import (
	"context"

	"github.com/ozline/grpc-todolist/cmd/user/model"
	"github.com/ozline/grpc-todolist/idl/pb/user"
)

func CreateUser(ctx context.Context, req *user.RegisterRequest) error {

	// err := DB.Where("username = ?", req.Username).First(&model.User{}).Error

	// if err == nil {
	// 	return errno.ErrUsernameAlreadyExists
	// }

	return DB.Create(&model.User{
		ID:       SF.NextVal(),
		Username: req.Username,
		Password: req.Password,
	}).Error
}

func GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("id = ?", id).First(user).Error
	return user, err
}

func GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	err := DB.Where("username = ?", username).First(user).Error
	return user, err
}

func CheckUserPassword(ctx context.Context, username string, password string) error {
	return DB.Where("username = ? AND password = ?", username, password).First(&model.User{}).Error
}
