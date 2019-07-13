package notificacaousuario
import (
	db "squad-3-aceleradev-fs-florianopolis/db"
	"squad-3-aceleradev-fs-florianopolis/entity"
	"strconv"
)

//Insert new notificacaousuario
func Insert(noteuser *entity.NotificacaoUsuario) error {
	erro := db.ExecutaComando("insert into notificacaousuario (idusuario, idnotificacao, data) values(" 
	+ strconv.FormatInt(noteuser.IDUsuario) + ",'" + strconv.FormatInt(noteuser.IDNotificacao) + "', '" 
	+ noteuser.Data.Format(time.RFC3339) + ");")
	return erro
}

//Get notificacaousuario
func Get(id int) (*entity.NotificacaoUsuario, error) {
	db, erro := db.Conect()
	seleciona, erro := db.Query("select * from notificacaousuario where id = " + strconv.Itoa(id))
	var noteuser entity.NotificacaoUsuario
	if erro == nil {

		for seleciona.Next() {
			erro := seleciona.Scan(&noteuser.IDUsuario, &noteuser.IDNotificacao, &noteuser.Data)
			if erro != nil {
				panic(erro.Error())
			}
		}
	}
	defer db.Close()
	return &noteuser, erro
}
