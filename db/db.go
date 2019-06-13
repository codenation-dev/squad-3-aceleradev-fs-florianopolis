package database
import (
	"strconv"	
	"database/sql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "bancouati_user"
	password = "1234567890"
	dbname   = "bancouati"
)


func Conect() (*sql.DB, error){
	db, erro := sql.Open("mysql", //"bancouati_user:1234567890@tcp(localhost:3306)/bancouati")
						 user + ":" + 
						 password + "@tcp(" + 
						 host + ":" + 
						 strconv.Itoa(port) + ")/" + 
						 dbname)
	return db, erro
}

func ExecutaComando(comando string) error{
	db, erro := Conect()
	retorno, erro := db.Query(comando);
	retorno.Close()
	return erro
}

