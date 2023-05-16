package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ozline/grpc-todolist/cmd/api/handler"
	"github.com/ozline/grpc-todolist/cmd/api/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	user := r.Group("/user")
	{
		user.POST("/register", handler.UserRegister)
		user.POST("/login", handler.UserLogin)
	}

	task := r.Group("/task").Use(middleware.JWT)
	{
		task.POST("/create", handler.TaskCreate)
		task.GET("/list", handler.TaskList)
		task.POST("/update", handler.TaskUpdate)
		// task.DELETE("/delete", handler.TaskDelete)
		task.POST("/delete", handler.TaskDelete)
	}

	return r
}
