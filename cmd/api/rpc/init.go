package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/ozline/grpc-todolist/idl/pb/task"
	"github.com/ozline/grpc-todolist/idl/pb/user"
	"github.com/ozline/grpc-todolist/pkg/discovery"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	UserClient user.UserServiceClient
	TaskClient task.TaskServiceClient
)

func Init() {
	Register = discovery.NewResolver([]string{viper.GetString("etcd.addr")}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()

	initUserClient()
	initTaskClient()
}

func Connect(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}

func initUserClient() {
	serviceName := viper.GetString("services.user.name")

	connUser, err := Connect(serviceName)

	if err != nil {
		panic(err)
	}

	UserClient = user.NewUserServiceClient(connUser)
}

func initTaskClient() {
	serviceName := viper.GetString("services.task.name")

	connTask, err := Connect(serviceName)

	if err != nil {
		panic(err)
	}

	TaskClient = task.NewTaskServiceClient(connTask)
}
