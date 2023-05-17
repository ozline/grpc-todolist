package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ozline/grpc-todolist/cmd/api/rpc"
	"github.com/ozline/grpc-todolist/idl/pb/experimental"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

func StreamPing(c *gin.Context) {

	resp, err := rpc.Ping(c, &experimental.Request{
		Ping: "ping",
	})

	if err != nil {
		BuildFailResponse(c, errno.NewErrNo(errno.ServiceInternalError.ErrorCode, err.Error()))
		return
	}

	BuildSuccessResponse(c, resp)
}

func ClientStream(c *gin.Context) {
	resp, err := rpc.ClientStream(c)

	if err != nil {
		BuildFailResponse(c, errno.NewErrNo(errno.ServiceInternalError.ErrorCode, err.Error()))
		return
	}

	BuildSuccessResponse(c, resp)
}

func ServerStream(c *gin.Context) {
	resp, err := rpc.ServerStream(c)

	if err != nil {
		BuildFailResponse(c, errno.NewErrNo(errno.ServiceInternalError.ErrorCode, err.Error()))
		return
	}

	BuildSuccessResponse(c, resp)
}

func BidirectionalStream(c *gin.Context) {
	resp, err := rpc.BidirectionalStream(c)

	if err != nil {
		BuildFailResponse(c, errno.NewErrNo(errno.ServiceInternalError.ErrorCode, err.Error()))
		return
	}

	BuildSuccessResponse(c, resp)
}
