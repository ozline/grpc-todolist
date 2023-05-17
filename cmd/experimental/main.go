package main

import (
	"flag"
	"log"
	"net"

	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/discovery"
	"github.com/ozline/grpc-todolist/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ozline/grpc-todolist/idl/pb/experimental"
)

func Init() *discovery.Register {
	// Args
	path := flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path, srvname)

	// etcd
	register := discovery.NewRegister([]string{config.Etcd.Addr}, logrus.New())
	// defer register.Stop()

	node := discovery.Server{
		Name: config.Service.Name,
		Addr: config.Service.Addr,
	}

	if _, err := register.Register(node, 10); err != nil {
		log.Fatalf("register service %s failed, err: %v", node.Name, err)
	}

	return register
}

func main() {
	register := Init()

	lis, err := net.Listen("tcp", config.Service.Addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	experimental.RegisterExperimentalServiceServer(s, NewExperimentalServiceImpl())
	reflection.Register(s) // Support server reflection

	go utils.ListenSignal(func() {
		register.Stop()
		s.Stop()
	})

	log.Printf("%s listening at %v\n", srvname, lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}