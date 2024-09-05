// db/database.go
package db

import (
    "os"                   // Importa o pacote os para acessar variáveis de ambiente
    "gorm.io/driver/mysql" // Importa o driver MySQL para GORM, o ORM utilizado
    "gorm.io/gorm"         // Importa o GORM, que é o ORM utilizado para interagir com o banco de dados
    "log"
)

// DB é uma variável global que armazena a conexão com o banco de dados.
// Ela é utilizada em várias partes do código para realizar operações no banco de dados.
var DB *gorm.DB

// ConnectDatabase conecta ao banco de dados MySQL.
// A função é responsável por abrir uma conexão com o banco de dados usando o GORM e armazená-la na variável DB.
// Esta função deve ser chamada no início da aplicação para garantir que o banco de dados esteja acessível.
func ConnectDatabase() {
    // Define o Data Source Name (DSN) que especifica as credenciais e o endereço do banco de dados MySQL.
    dsn := os.Getenv("MYSQL_DSN")
    
    // Abre a conexão com o banco de dados usando o driver MySQL.
    // mysql.Open(dsn) inicializa o driver MySQL para o GORM.
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    
    // Se ocorrer um erro ao conectar, a aplicação é encerrada e um log de erro é exibido.
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }

    // AutoMigrate é usado para criar ou atualizar o esquema do banco de dados para o modelo User.
    // Em TDD, você frequentemente escreve testes para verificar se a migração de banco de dados ocorre corretamente.
    database.AutoMigrate(&User{})

    // A conexão de banco de dados é armazenada na variável global DB.
    DB = database
}

// User representa um modelo de usuário simples.
// Este struct é usado pelo GORM para mapear a tabela 'users' no banco de dados.
// No TDD, escreveríamos testes para garantir que as operações CRUD funcionam conforme esperado para este modelo.
type User struct {
    ID   uint   // ID é a chave primária do usuário.
    Name string // Name é um campo simples de texto para armazenar o nome do usuário.
}
