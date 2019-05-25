package main

import(
	"squad-3-aceleradev-fs-florianopolis/model"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	//"squad-3-aceleradev-fs-florianopolis/entity"
)

func main(){
	usuario, erro := model.GetUsuario(1)
	//newuser := entity.Usuario{ID:0,Cpf:99999999999,Nome:"Norberto Fonseca",Senha:"1234",Email:"o@o.com",FuncionarioPuplico:false}
	//erro = model.Insert(&newuser)
	//erro = model.Delete(6)
	if erro != nil{
		panic(erro.Error())
	}
	fmt.Println(usuario)
}
