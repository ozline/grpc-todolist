package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ozline/grpc-todolist/cmd/api/types"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SendResponse(c *gin.Context, base types.BaseResp, data interface{}) {
	c.JSON(200, gin.H{
		"code": base.Code,
		"msg":  base.Msg,
		"data": data,
	})
}

func BuildBaseResp(code int64, msg string) types.BaseResp {
	return types.BaseResp{Code: code, Msg: msg}
}

func BuildSuccessResponse(c *gin.Context, data interface{}) {
	SendResponse(c, types.BaseResp{Code: errno.SuccessCode, Msg: errno.SuccessMsg}, data)
}

func BuildFailResponse(c *gin.Context, err error) {
	errno := errno.ConvertErr(err)
	SendResponse(c, types.BaseResp{Code: errno.ErrorCode, Msg: errno.ErrorMsg}, nil)
}

func BuildFailResponseWithBaseResp(c *gin.Context, base types.BaseResp) {
	SendResponse(c, base, nil)
}
