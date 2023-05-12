package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ozline/grpc-todolist/cmd/api/rpc"
	"github.com/ozline/grpc-todolist/cmd/api/types"
	"github.com/ozline/grpc-todolist/idl/pb/task"
	"github.com/ozline/grpc-todolist/pkg/errno"
)

func TaskCreate(c *gin.Context) {
	var req types.TaskCreateRequest

	err := c.Bind(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	resp, err := rpc.TaskCreate(c, &task.CreateRequest{
		UserId:  c.GetInt64("userid"),
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	BuildSuccessResponse(c, resp)
}

func TaskList(c *gin.Context) {
	var req types.TaskListRequest

	err := c.BindQuery(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	resp, err := rpc.TaskGetList(c, &task.GetListRequest{
		UserId: c.GetInt64("userid"),
		Status: *req.Status,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	BuildSuccessResponse(c, resp)
}

func TaskDelete(c *gin.Context) {
	var req types.TaskDeleteRequest

	err := c.Bind(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	err = rpc.TaskDelete(c, &task.DeleteRequest{
		Id: req.ID,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	BuildSuccessResponse(c, nil)
}

func TaskUpdate(c *gin.Context) {
	var req types.TaskUpdateRequest

	err := c.Bind(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	resp, err := rpc.TaskUpdate(c, &task.UpdateRequest{
		Id:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	BuildSuccessResponse(c, resp)
}
