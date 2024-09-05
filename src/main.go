package main

import (
    "src/handlers"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Configurar as rotas
    router.GET("/users", handlers.GetUsers)
    router.POST("/users", handlers.CreateUser)

    // Iniciar o servidor
    http.ListenAndServe(":8080", router)
}
