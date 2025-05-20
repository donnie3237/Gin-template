package main

import (
    "myapp/config"
    "myapp/internal/handler"
    "myapp/internal/repository"
    "myapp/internal/route"
    "myapp/internal/service"
)

func main() {
    config.LoadEnv()

    db := config.ConnectMongoDB()
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    r := route.SetupRouter(userHandler)
    r.Run(":8080")
}
