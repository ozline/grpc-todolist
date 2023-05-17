// 这个文件定义了一个实验性的服务处理器，其中包含了一个实验性的服务实现。
// 这个服务实现包含了四个方法，分别是 Ping、ClientStream、ServerStream 和 BidirectionalStream。
// 这四个方法分别实现了不同类型的 gRPC 流式调用。
// 具体来说，Ping 方法是一个普通的 Unary 调用，返回一个响应；
// ClientStream 方法是一个客户端流式调用，允许客户端发送最多 6 个消息，然后服务端返回一个响应；
// ServerStream 方法是一个服务端流式调用，允许客户端发送一个消息，然后服务端返回最多 10 个响应；
// BidirectionalStream 方法是一个双向流式调用，允许客户端发送任意多个消息，每 2 个消息服务端将会响应一个消息。
// 这个服务处理器是基于 gRPC 的，使用了 idl/pb/experimental 包中定义的实验性服务接口。

// Gernerated by chatGPT 4.0

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

// 普通Unary调用
// 服务端返回一个响应
func (es *ExperimentalServiceImpl) Ping(ctx context.Context, req *experimental.Request) (resp *experimental.Response, err error) {
	resp = &experimental.Response{
		Pong: "pong",
	}

	return
}

// 客户端流式推送到服务端
// 我们允许客户端发送最多6个消息，然后服务端返回一个响应
func (es *ExperimentalServiceImpl) ClientStream(stream experimental.ExperimentalService_ClientStreamServer) error {

	msgs := []string{}

	for {
		if len(msgs) > 6 {
			return stream.SendAndClose(&experimental.Response{Pong: "Maxnum pings is 6, enough!"})
		}

		msg, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&experimental.Response{Pong: fmt.Sprintf("Received %d requests", len(msgs))})
			}
			log.Printf("receive msg error: %v", err)
			return err
		}

		msgs = append(msgs, msg.Ping)
	}
}

// 服务端流式响应客户端
// 我们允许客户端发送一个消息，然后服务端返回最多10个响应
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

// 双向流式通信
// 我们允许客户端发送任意多个消息，每2个消息服务端将会响应一个消息
func (es *ExperimentalServiceImpl) BidirectionalStream(stream experimental.ExperimentalService_BidirectionalStreamServer) error {

	msgs := []string{}
	for {
		msg, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("receive msg error: %v", err)
			return err
		}

		msgs = append(msgs, msg.Ping)

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
