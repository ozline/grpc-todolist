package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/idl/pb/experimental"
	"github.com/ozline/grpc-todolist/idl/pb/task"
	"github.com/ozline/grpc-todolist/idl/pb/user"
	"github.com/ozline/grpc-todolist/pkg/discovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	ExperimentalClient experimental.ExperimentalServiceClient
)

const (
	experimentalSrvName = "experimental"
)

func Init() {
	Register = discovery.NewResolver([]string{config.Etcd.Addr}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()

	initClient(config.GetService(experimentalSrvName).Name, &ExperimentalClient)
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
	case *experimental.ExperimentalServiceClient:
		*c = experimental.NewExperimentalServiceClient(conn)
	default:
		panic("unsupported client type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	// Load balance
	if config.GetService(serviceName).LB {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}

	return grpc.DialContext(ctx, addr, opts...)
}
