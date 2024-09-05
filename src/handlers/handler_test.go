package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/gin-gonic/gin"
    "src/db" // Importa o pacote db para conectar ao banco de dados
)

// TestMain é a função principal para configuração de testes
func TestMain(m *testing.M) {
    // Defina a variável de ambiente MYSQL_DSN antes de conectar ao banco de dados
    os.Setenv("MYSQL_DSN", "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")

    // Conecte ao banco de dados
    db.ConnectDatabase()

    // Execute os testes
    m.Run()
}

// TestGetUsers verifica a funcionalidade de obter usuários
func TestGetUsers(t *testing.T) {
    // Inicializa o Gin para testes
    gin.SetMode(gin.TestMode)
    r := gin.Default()

    // Roteia o endpoint para o handler GetUsers
    r.GET("/users", GetUsers)

    // Cria uma requisição de teste GET para o endpoint "/users"
    req, _ := http.NewRequest(http.MethodGet, "/users", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Verifica se o status retornado é 200 OK
    if w.Code != http.StatusOK {
        t.Errorf("expected status 200 but got %d", w.Code)
    }
}

// TestCreateUser verifica a funcionalidade de criar um usuário
func TestCreateUser(t *testing.T) {
    // Inicializa o Gin para testes
    gin.SetMode(gin.TestMode)
    r := gin.Default()

    // Roteia o endpoint para o handler CreateUser
    r.POST("/users", CreateUser)

    // Cria um corpo de requisição JSON para o teste
    body := `{"Name": "Vitorhrds"}`
    req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(body)))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    // Verifica se o status retornado é 200 OK
    if w.Code != http.StatusOK {
        t.Errorf("expected status 200 but got %d", w.Code)
    }

    // Verifica se o nome retornado contém "John Doe"
    if !bytes.Contains(w.Body.Bytes(), []byte("Vitorhrds")) {
        t.Errorf("response body does not contain 'Vitorhrds'")
    }
}
