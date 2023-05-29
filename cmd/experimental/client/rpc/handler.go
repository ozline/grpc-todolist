package rpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/ozline/grpc-todolist/idl/pb/experimental"
)

func Ping(ctx context.Context, req *experimental.Request) (string, error) {
	resp, err := ExperimentalClient.Ping(ctx, req)

	if err != nil {
		return "", err
	}

	return resp.Pong, nil
}

func ClientStream(ctx context.Context) (string, error) {

	stream, err := ExperimentalClient.ClientStream(ctx)

	if err != nil {
		return "", err
	}

	// sendMsg
	for i := 0; i < 5; i++ {
		err = stream.Send(&experimental.Request{Ping: fmt.Sprintf("the %dth ping", i)})

		if err != nil {
			return "", err
		}
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		return "", err
	}

	return res.Pong, nil
}

func ServerStream(ctx context.Context) ([]string, error) {

	stream, err := ExperimentalClient.ServerStream(ctx, &experimental.Request{
		Ping: "114514",
	})

	if err != nil {
		return nil, err
	}

	// sendMsg

	resp := make([]string, 0)

	for {
		msg, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		resp = append(resp, msg.Pong)
	}

	return resp, nil
}

func BidirectionalStream(ctx context.Context) ([]string, error) {
	stream, err := ExperimentalClient.BidirectionalStream(ctx)

	if err != nil {
		return nil, err
	}

	c := ReceiveStream(stream)

	// sendMsg
	go func() {
		for i := 0; i < 6; i++ {
			err = stream.Send(&experimental.Request{Ping: "114514"})
			if err != nil {
				return
			}
		}
		err = stream.CloseSend()
		if err != nil {
			return
		}
	}()

	var resp []string
	for msg := range c {
		if msg == nil {
			return nil, errors.New("error receiving stream")
		}
		resp = append(resp, msg...)
	}

	return resp, nil
}

func ReceiveStream(stream experimental.ExperimentalService_BidirectionalStreamClient) chan []string {
	resp := make(chan []string)
	go func() {
		defer close(resp)

		for {
			msg, err := stream.Recv()

			if err != nil {
				if err == io.EOF {
					log.Printf("stream closed")
					break
				}
				log.Printf("error receiving stream: %v", err)
				resp <- nil
				return
			}
			resp <- []string{msg.Pong}
		}
	}()

	return resp
}
