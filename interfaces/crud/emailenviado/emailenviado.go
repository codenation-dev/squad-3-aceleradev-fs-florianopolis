package emailenviado

import (
	"strconv"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	utils "squad-3-aceleradev-fs-florianopolis/utils"
	"time"
)

//Insert into emailenviado
func Insert(ID int,Endereco string) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Insert(EMAILENVIADO)", erro.Error())
	}
	defer dbi.Database.Close()
	t := time.Now()
	formatTime := t.Format("2006-01-02 15:04:05")
	logs.Info("Insert(EMAILENVIADO)", "Inserting new EMAILENVIADO in DB...")
	_, err := dbi.Database.Query(`INSERT INTO EMAILENVIADO (idnotificacao, emailusuario, data) VALUES(?, ?, ?)`, 
								 ID, Endereco, formatTime)
	return err
}

//GetAll gets a list
func GetAll(id int) ([]entity.EmailEnviado) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("get(EMAILENVIADO)",erro.Error())
	}
	defer dbi.Database.Close()
	seleciona, erro := dbi.Database.Query(`SELECT ID, EmailUsuario, Data FROM EMAILENVIADO WHERE idnotificacao = ` + strconv.Itoa(id))
	var note entity.EmailEnviado
	var List []entity.EmailEnviado
	if erro == nil {
		for seleciona.Next() {
			var d string
			erro := seleciona.Scan(&note.ID, &note.EmailUsuario, &d)
			note.Data = utils.ConvertDateTimeSQL(d)
			if erro != nil {
				logs.Errorf("Get(notificação)", erro.Error())
			}
			List = append(List,note)
		}
	}
	return List
}