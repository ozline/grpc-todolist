package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/ozline/grpc-todolist/cmd/api/routes"
	"github.com/ozline/grpc-todolist/cmd/api/rpc"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/utils"
	"github.com/spf13/viper"
)

func Init() {
	path := flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path)

	rpc.Init()
}

func main() {
	Init()

	r := routes.NewRouter()

	server := &http.Server{
		Addr:           viper.GetString("services.api.addr"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go utils.ListenSignal(func() {
		server.Shutdown(context.TODO())
	})

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
