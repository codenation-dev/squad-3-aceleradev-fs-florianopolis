package database

import (
	"os"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"

	//"github.com/codenation-dev/squad-3-aceleradev-fs-florianopolis/entities/logs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var hostAddress = os.Getenv("MYSQL_HOST")
var portAddress = os.Getenv("MYSQL_PORT")
var userDB = os.Getenv("MYSQL_USER")
var passwordDB = os.Getenv("MYSQL_PASSWORD")
var databaseName = os.Getenv("MYSQL_DATABASE")

/*var hostAddress = "localhost"
var portAddress = "3306"
var userDB = "root"
var passwordDB = ""
var databaseName = "PROJETOUATI"*/

// MySQLDatabase struct for db
type MySQLDatabase struct {
	Database *sql.DB
}

// Init the MySQL DB and return the struct reference
func Init() (*MySQLDatabase, error) {
	fmt.Println("CONEXÃO...")
	//db, err := sql.Open("mysql", "user:password@/dbname")
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userDB, passwordDB, hostAddress, portAddress, databaseName)
	dbConnection, erro := sql.Open("mysql", ConnectionString)

	if erro != nil {
		logs.Errorf("MySQL Database", fmt.Sprintf("%s", erro))
		return nil, erro
	}

	MyDB := &MySQLDatabase{Database: dbConnection}

	return MyDB, nil
}

//ExecQuery Execute a Query
//Alterado para retornar o resultado da query (junto com o erro), para casos de Select
func (MyDB MySQLDatabase) ExecQuery(comando string) (error, *sql.Rows) {
	retorno, erro := MyDB.Database.Query(comando);if erro != nil {
		logs.Errorf("MySQL Database", fmt.Sprintf("%s", erro))
		fmt.Println("Erro (ExecQuery): ", erro)
	}
	defer MyDB.Database.Close()
	return erro, retorno
}
