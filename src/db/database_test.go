// db/database_test.go
package db

import (
    "os"
    "testing" // Pacote testing é usado para escrever testes unitários em Go.

    "github.com/stretchr/testify/suite" // Importa testify, uma biblioteca de asserts e suites de testes.
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// TestMain é a função principal para configuração de testes
func TestMain(m *testing.M) {
    // Defina a variável de ambiente MYSQL_DSN antes de conectar ao banco de dados
    os.Setenv("MYSQL_DSN", "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")

    // Conecte ao banco de dados
    ConnectDatabase()

    // Execute os testes
    m.Run()
}

// DatabaseTestSuite é uma suite de testes que agrupa testes relacionados ao banco de dados.
// Usar suites é uma prática recomendada para organizar e estruturar testes de maneira lógica.
type DatabaseTestSuite struct {
    suite.Suite // Embeds testify's suite functionality
    db *gorm.DB // Armazena a conexão de banco de dados para uso em testes.
}

// SetupSuite é executado antes de todos os testes na suite.
// Em TDD, isso é usado para preparar o estado inicial do teste, garantindo que todos os testes sejam independentes e repetíveis.
func (suite *DatabaseTestSuite) SetupSuite() {
    // Define o DSN para o banco de dados de teste.
    dsn := os.Getenv("MYSQL_DSN")
    
    // Abre a conexão com o banco de dados de teste.
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    suite.Require().NoError(err, "Erro ao conectar ao banco de dados de teste")

    suite.db = db.Debug() // Habilita o modo debug para visualizar todas as operações SQL.
    
    // AutoMigrate é usado aqui para garantir que a tabela do modelo User seja criada no banco de dados de teste.
    err = suite.db.AutoMigrate(&User{})
    suite.Require().NoError(err, "Erro ao migrar tabelas do banco de dados")
}

// TestUserInsertion testa a inserção de um registro de usuário no banco de dados.
// Em TDD, primeiro escrevemos este teste para definir o comportamento esperado de inserção.
func (suite *DatabaseTestSuite) TestUserInsertion() {
    user := User{Name: "John Doe"} // Cria um novo objeto User.
    
    // Tenta inserir o objeto User no banco de dados.
    err := suite.db.Create(&user).Error
    suite.Require().NoError(err, "Erro ao criar registro de usuário")

    var retrievedUser User
    // Tenta recuperar o usuário recém-criado do banco de dados.
    err = suite.db.First(&retrievedUser, "name = ?", "John Doe").Error
    suite.Require().NoError(err, "Erro ao recuperar registro de usuário")

    // Compara o nome do usuário inserido com o nome recuperado para garantir que são iguais.
    suite.Equal(user.Name, retrievedUser.Name, "Os nomes devem coincidir")
}

// TearDownSuite é executado após todos os testes na suite.
// Limpa o banco de dados de teste para garantir que os testes futuros não sejam afetados.
func (suite *DatabaseTestSuite) TearDownSuite() {
    err := suite.db.Exec("DROP TABLE users;").Error // Remove a tabela de teste do banco de dados.
    suite.Require().NoError(err, "Erro ao remover a tabela de teste")

    sqlDB, _ := suite.db.DB() // Recupera a conexão SQL subjacente para fechá-la.
    err = sqlDB.Close()
    suite.Require().NoError(err, "Erro ao fechar o banco de dados de teste")
}

// TestDatabaseTestSuite é o ponto de entrada para rodar a suite de testes.
// Ele verifica uma variável de ambiente para decidir se os testes de banco de dados devem ser executados.
func TestDatabaseTestSuite(t *testing.T) {
    if os.Getenv("MYSQL_DSN") == "" {
        t.Skip("Pulando os testes de MySQL; forneça a variável de ambiente MYSQL_DSN")
    }

    suite.Run(t, new(DatabaseTestSuite))
}
