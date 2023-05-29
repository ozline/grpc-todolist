package main

import (
	"flag"
	"fmt"

	"github.com/ozline/grpc-todolist/cmd/experimental/client/rpc"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/idl/pb/experimental"
	"golang.org/x/net/context"
)

var (
	path    *string
	srvnum  *int
	srvname = "api"
)

func main() {
	path = flag.String("config", "./config", "config path")
	srvnum = flag.Int("srvnum", 0, "node number")
	flag.Parse()
	config.Init(*path, srvname, *srvnum)

	rpc.Init()

	c := context.Background()

	// ping
	resp, err := rpc.Ping(c, &experimental.Request{
		Ping: "ping",
	})

	if err != nil {
		fmt.Println(err)
	}

	println(resp)

	// client stream
	resp, err = rpc.ClientStream(c)

	if err != nil {
		fmt.Println(err)
	}

	println(resp)

	// server stream
	resplist, err := rpc.ServerStream(c)

	if err != nil {
		fmt.Println(err)
	}

	for _, resp := range resplist {
		println(resp)
	}

	// bidirectional stream
	resplist, err = rpc.BidirectionalStream(c)

	if err != nil {
		fmt.Println(err)
	}

	for _, resp := range resplist {
		println(resp)
	}

	fmt.Println("done")
}
