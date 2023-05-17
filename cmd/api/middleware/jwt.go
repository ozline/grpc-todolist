package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozline/grpc-todolist/cmd/api/handler"
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/errno"
	"github.com/ozline/grpc-todolist/pkg/utils"
)

const (
	authHeader = "Authorization"
	key        = "userid"
)

func JWT(c *gin.Context) {
	var respErr = errno.Success
	var claims *utils.Claims
	var err error

	token := c.GetHeader(authHeader)
	if token == "" {
		respErr = errno.AuthorizationFailError
	} else {
		claims, err = utils.ParseToken(token, config.Server.Secret)
		if err != nil {
			respErr = errno.AuthorizationFailError
		} else if time.Now().Unix() > claims.ExpiresAt {
			respErr = errno.AuthorizationExpiredError
		}
	}

	if respErr != errno.Success {
		handler.BuildFailResponse(c, respErr)
		c.Abort()
		return
	}

	token, err = utils.GenerateToken(claims.UserID, config.Server.Secret)

	if err != nil {
		handler.BuildFailResponse(c, errno.AuthorizationFailError.WithMessage(err.Error()))
		c.Abort()
		return
	}
	c.Header(authHeader, token)
	c.Set(key, claims.UserID)

	c.Next()
}
