package usuario
import(
	"strconv"	
	"squad-3-aceleradev-fs-florianopolis/entity"
	db "squad-3-aceleradev-fs-florianopolis/db"
)
//Insert new usuario
func Insert (user *entity.Usuario) error{	
	erro := db.ExecutaComando("insert into usuario (cpf, nome, senha, email, funcionariopublico) values(" + strconv.FormatInt(user.Cpf, 10) + ",'" + user.Nome + "', '" + user.Senha + "','" + user.Email + "',"+ strconv.FormatBool(user.FuncionarioPuplico) +");")
	return erro
}
//Delete usuario by id
func Delete(id int) error {
	erro := db.ExecutaComando("delete from usuario where id = " + strconv.Itoa(id))
	return erro
}
//Update usuario
func Update (user *entity.Usuario) error{
	erro := db.ExecutaComando("update usuario set cpf = "+strconv.FormatInt(user.Cpf, 10) + ", nome = '" + user.Nome + "', senha = '" + user.Senha + "', email = '" + user.Email + "', funcionariopublico = " + strconv.FormatBool(user.FuncionarioPuplico) + " where id = " + strconv.Itoa(user.ID))
	return erro
}
//Get usuario by id
func Get(id int) (*entity.Usuario, error){
	db, erro := db.Conect()
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