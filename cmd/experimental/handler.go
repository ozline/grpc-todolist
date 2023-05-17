package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ozline/grpc-todolist/idl/pb/experimental"
)

type ExperimentalServiceImpl struct {
	experimental.UnimplementedExperimentalServiceServer
}

func NewExperimentalServiceImpl() *ExperimentalServiceImpl {
	return &ExperimentalServiceImpl{}
}

func (es *ExperimentalServiceImpl) Ping(ctx context.Context, req *experimental.Request) (resp *experimental.Response, err error) {
	resp = &experimental.Response{
		Pong: "pong",
	}

	return
}

func (es *ExperimentalServiceImpl) ClientStream(stream experimental.ExperimentalService_ClientStreamServer) error {

	msgs := []string{}

	for {
		// 提前结束接收消息
		if len(msgs) > 10 {
			return stream.SendAndClose(&experimental.Response{Pong: "Maxnum pings is 6, enough!"})
		}

		msg, err := stream.Recv()

		if err != nil {
			// 客户端消息结束
			if err == io.EOF {
				return stream.SendAndClose(&experimental.Response{Pong: fmt.Sprintf("Received %d requests", len(msgs))})
			}
			log.Printf("receive msg error: %v", err)
			return err
		}

		msgs = append(msgs, msg.Ping)
		// stream.SendMsg(&experimental.Response{Pong: fmt.Sprintf("the %dth '%s'", len(msgs), msg.Ping)})
	}
}

func (es *ExperimentalServiceImpl) ServerStream(req *experimental.Request, stream experimental.ExperimentalService_ServerStreamServer) error {
	for i := 0; i < 10; i++ {
		data := &experimental.Response{Pong: fmt.Sprintf("the %dth pong", i)}

		err := stream.Send(data)

		if err != nil {
			return err
		}
	}
	return nil
}

func (es *ExperimentalServiceImpl) BidirectionalStream(stream experimental.ExperimentalService_BidirectionalStreamServer) error {

	msgs := []string{}
	for {
		// receive msg
		msg, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("receive msg error: %v", err)
			return err
		}

		msgs = append(msgs, msg.Ping)

		// 每收到2个消息响应一次
		if len(msgs)%2 == 0 {
			err = stream.SendMsg(&experimental.Response{
				Pong: "pong - len of msgs is " + fmt.Sprintf("%d", len(msgs)),
			})

			if err != nil {
				log.Printf("receive msg error: %v", err)
				return err
			}
		}
	}
	return nil
}
