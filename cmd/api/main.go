package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/ozline/grpc-todolist/cmd/api/routes"
	"github.com/ozline/grpc-todolist/cmd/api/rpc"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/utils"
)

var (
	path   *string
	srvnum *int
)

func Init() {
	path = flag.String("config", "./config", "config path")
	srvnum = flag.Int("srvnum", 0, "node number")
	flag.Parse()
	config.Init(*path, srvname, *srvnum)

	rpc.Init()
}

func main() {
	Init()

	r := routes.NewRouter()

	server := &http.Server{
		Addr:           config.Service.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go utils.ListenSignal(func() {
		server.Shutdown(context.TODO())
	})

	log.Printf("%s listening at %v\n", srvname, server.Addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
