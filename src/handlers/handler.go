package handlers

import (
    "net/http"
    "src/db"
    "os"

    "github.com/gin-gonic/gin"
)

func init() {
    // Defina a variável de ambiente MYSQL_DSN antes de conectar ao banco de dados
    os.Setenv("MYSQL_DSN", "vitorhrds:vitor12345@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")

    // Conecte ao banco de dados
    db.ConnectDatabase()
}

// GetUsers manipula a requisição para listar todos os usuários
func GetUsers(c *gin.Context) {
    var users []db.User

    // Certifique-se de que a conexão de banco de dados foi inicializada
    if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
        return
    }

    // Recupera todos os usuários do banco de dados
    if err := db.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

// CreateUser manipula a requisição para criar um novo usuário
func CreateUser(c *gin.Context) {
    var user db.User

    // Certifique-se de que a conexão de banco de dados foi inicializada
    if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
        return
    }

    // Faz o binding dos dados da requisição JSON para o objeto user
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insere o novo usuário no banco de dados
    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}
