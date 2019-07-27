package database

import (
	"database/sql"
	"fmt"
	"os"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"

	_ "github.com/go-sql-driver/mysql"
)

var hostAddress = os.Getenv("MYSQL_HOST")
var portAddress = os.Getenv("MYSQL_PORT")
var userDB = os.Getenv("MYSQL_USER")
var passwordDB = os.Getenv("MYSQL_PASSWORD")
var databaseName = os.Getenv("MYSQL_DATABASE")

// MySQLDatabase struct for db
type MySQLDatabase struct {
	Database *sql.DB
}

// Init the MySQL DB and return the struct reference
func Init() (*MySQLDatabase, error) {

	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userDB, passwordDB, hostAddress, portAddress, databaseName)
	dbConnection, erro := sql.Open("mysql", ConnectionString)
	if erro != nil {
		logs.Errorf("MySQL Database", fmt.Sprintf("%s", erro))
		return nil, erro
	}
	MyDB := &MySQLDatabase{Database: dbConnection}
	return MyDB, nil
}