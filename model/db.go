package model
import (
	"squad-3-aceleradev-fs-florianopolis/entity"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "pass"
	dbname   = "dbname"
)


func conect() (*sql.DB, error){
	db, erro := sql.Open("mysql", 
						 user + ":" + 
						 password + "@tcp(" + 
						 host + ":" + 
						 string(port) + ")/" + 
						 dbname)
	if erro != nil{
		panic(erro.Error())
	}
	defer db.Close()
	return db, erro
}

func Insert (user *entity.Usuario) error{
	db, erro := conect()
	insere, erro := db.Query("insert into usuario values(" + 
								string(user.Cpf) +","+
								"'" + user.Nome + "'," +
								"'" + user.Senha + "'," +
								"'" + user.Email + "'," +
								"'" + user.FuncionarioPuplico +"');");
	if erro != nil{
		panic(erro.Error())
	}
	insere.Close()
	return erro
}

func GetUsuario(id int) (*entity.Usuario, error){
	db, erro := conect()
	seleciona, erro := db.Query("select * from usuario where id = " + string(id))
	
	if erro != nil{
		panic(erro.Error())
	}

	var user  *entity.Usuario
	for seleciona.Next() {
		erro := seleciona.Scan(&user.Cpf, &user.Nome, &user.Senha, &user.Email, &user.FuncionarioPuplico)	
		if erro != nil{
			panic(erro.Error())
		}
	}
	
	return user, erro
}