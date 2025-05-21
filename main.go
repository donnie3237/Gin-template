package main

import (
    "mytemplate/config"
    "mytemplate/internal/handler"
    "mytemplate/internal/repository"
    "mytemplate/internal/router"
    "mytemplate/internal/service"
    "fmt"
)

func main() {
    config.LoadEnv()

    db := config.ConnectMongoDB()
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    r := router.SetupRouter(userHandler)
    r.Run(":8083")
    fmt.Println("Hello, mytemplate!")
}
