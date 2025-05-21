package router

import (
    "github.com/gin-gonic/gin"
    "mytemplate/internal/handler"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
    r := gin.Default()
    r.POST("/users", userHandler.CreateUser)
    return r
}
