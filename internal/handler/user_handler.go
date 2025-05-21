package handler

import (
    "context"
    "net/http"
    "time"

    "mytemplate/internal/model"
    "mytemplate/internal/service"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
    Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = primitive.NewObjectID()
    user.CreatedAt = time.Now().Unix()

    if err := h.Service.CreateUser(context.Background(), &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}
