package model
import (
	"strconv"
	"squad-3-aceleradev-fs-florianopolis/entity"
	"database/sql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "bancouati_user"
	password = "1234567890"
	dbname   = "bancouati"
)


func conect() (*sql.DB, error){
	db, erro := sql.Open("mysql", //"bancouati_user:1234567890@tcp(localhost:3306)/bancouati")
						 user + ":" + 
						 password + "@tcp(" + 
						 host + ":" + 
						 strconv.Itoa(port) + ")/" + 
						 dbname)
	return db, erro
}

func executaComando(comando string) error{
	db, erro := conect()
	retorno, erro := db.Query(comando);
	retorno.Close()
	return erro
}

func Insert (user *entity.Usuario) error{	
	erro := executaComando("insert into usuario (cpf, nome, senha, email, funcionariopublico) values(" + strconv.FormatInt(user.Cpf, 10) + ",'" + user.Nome + "', '" + user.Senha + "','" + user.Email + "',"+ strconv.FormatBool(user.FuncionarioPuplico) +");")
	return erro
}

func Delete(id int) error {
	erro := executaComando("delete from usuario where id = " + strconv.Itoa(id))
	return erro
}

func Update (user *entity.Usuario) error{
	erro := executaComando("update usuario set cpf = "+strconv.FormatInt(user.Cpf, 10) + ", nome = '" + user.Nome + "', senha = '" + user.Senha + "', email = '" + user.Email + "', funcionariopublico = " + strconv.FormatBool(user.FuncionarioPuplico) + " where id = " + strconv.Itoa(user.ID))
	return erro
}

func GetUsuario(id int) (*entity.Usuario, error){
	db, erro := conect()
	seleciona, erro := db.Query("select * from usuario where id = " +  strconv.Itoa(id))
	var user entity.Usuario
	if erro == nil{
		
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Cpf, &user.Nome, &user.Senha, &user.Email, &user.FuncionarioPuplico)	
			if erro != nil{
				panic(erro.Error())
			}
		}
	}
	defer db.Close()
	return &user, erro
}