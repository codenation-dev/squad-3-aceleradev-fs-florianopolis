package notificacao

import (
	db "squad-3-aceleradev-fs-florianopolis/db"
	"squad-3-aceleradev-fs-florianopolis/entity"
	"strconv"
)

//Insert new notificacao
func Insert(note *entity.Notificacao) error {
	erro := db.ExecutaComando("insert into notificacao (idpessoa) values(" + strconv.Itoa(note.IDPessoa) + ");")
	return erro
}

//Delete notificacao by ID
func Delete(id int) error {
	erro := db.ExecutaComando("delete from notificacao where id = " + strconv.Itoa(id))
	return erro
}

//Update notificacao
func Update(note *entity.Notificacao) error {
	erro := db.ExecutaComando("update notificacao set idpessoa = "  + strconv.Itoa(note.IDPessoa) 
	+ " where id = " + strconv.Itoa(note.ID))
	return erro
}

//Get notificacao by ID
func Get(id int) (*entity.Notificacao, error) {
	db, erro := db.Conect()
	seleciona, erro := db.Query("select * from notificacao where id = " + strconv.Itoa(id))
	var note entity.Notificacao
	if erro == nil {

		for seleciona.Next() {
			erro := seleciona.Scan(&note.ID, &note.IDPessoa)
			if erro != nil {
				panic(erro.Error())
			}
		}
	}
	defer db.Close()
	return &note, erro
}
