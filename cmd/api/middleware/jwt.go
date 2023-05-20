package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ozline/grpc-todolist/cmd/api/handler"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/consts"
	"github.com/ozline/grpc-todolist/pkg/errno"
	"github.com/ozline/grpc-todolist/pkg/utils"
)

func JWT(c *gin.Context) {
	var claims *utils.Claims
	var err error

	token := c.GetHeader(consts.AuthHeader)
	if token == "" {
		handler.BuildFailResponse(c, errno.AuthorizationFailError)
		c.Abort()
		return
	}

	claims, err = utils.ParseToken(token, config.Server.Secret)
	if err != nil {
		handler.BuildFailResponse(c, err)
		c.Abort()
		return
	}

	token, err = utils.GenerateToken(claims.UserID, config.Server.Secret)

	if err != nil {
		handler.BuildFailResponse(c, errno.NewErrNo(errno.AuthorizationFailedErrCode, err.Error()))
		c.Abort()
		return
	}
	c.Header(consts.AuthHeader, token)
	c.Set(consts.KeyUserId, claims.UserID)

	c.Next()
}
