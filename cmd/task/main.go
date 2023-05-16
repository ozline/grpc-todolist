package main

import (
	"flag"
	"log"
	"net"

	"github.com/ozline/grpc-todolist/cmd/task/dal"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/discovery"
	"github.com/ozline/grpc-todolist/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	service "github.com/ozline/grpc-todolist/idl/pb/task"
)

func Init() *discovery.Register {
	// Args
	path := flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path)

	// Dal
	dal.Init()

	// etcd
	register := discovery.NewRegister([]string{viper.GetString("etcd.addr")}, logrus.New())

	node := discovery.Server{
		Name: viper.GetString("services.task.name"),
		Addr: viper.GetString("services.task.addr"),
	}

	if _, err := register.Register(node, 10); err != nil {
		log.Fatalf("register service %s failed, err: %v", node.Name, err)
	}

	return register
}

func main() {
	register := Init()

	lis, err := net.Listen("tcp", viper.GetString("services.task.addr"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	service.RegisterTaskServiceServer(s, NewTaskServiceImpl())
	reflection.Register(s) // Support server reflection

	log.Printf("task listening at %v\n", lis.Addr())

	go utils.ListenSignal(func() {
		register.Stop()
		s.Stop()
	})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
