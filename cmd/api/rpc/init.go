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

	initClient(viper.GetString("services.user.name"), &UserClient)
	initClient(viper.GetString("services.task.name"), &TaskClient)

}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}

	switch c := client.(type) {
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *task.TaskServiceClient:
		*c = task.NewTaskServiceClient(conn)
	default:
		panic("unsupported client type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
