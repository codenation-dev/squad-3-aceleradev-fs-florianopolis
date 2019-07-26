package emailenviado

import (
	"strconv"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"time"
)

//Insert into emailenviado
func Insert(emailenviado *entity.EmailEnviado) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Insert(EMAILENVIADO)", erro.Error())
	}
	defer dbi.Database.Close()
	t := time.Now()
	formatTime := t.Format("2006-01-02 15:04:05")
	logs.Info("Insert(EMAILENVIADO)", "Inserting new EMAILENVIADO in DB...")
	_, err := dbi.Database.Query(`INSERT INTO EMAILENVIADO (idnotificacao, emailusuario, data) VALUES(?, ?, ?)`, 
								 emailenviado.IDNotificacao, emailenviado.EmailUsuario, formatTime)
	return err
}

func GetAll(pData time.Time) (*entity.EmailEnviado, error) {
	var Data time.Time
	var note entity.Notificacao
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("get(EMAILENVIADO)", erro.Error())
	}
	defer dbi.Database.Close()
	if pData != Data{
		formatTime := pData.Format("2006-01-02 15:04:05")	
		_, err := dbi.Database.Query(`select * from EMAILENVIADO order by data desc limit 1`)
	}else {		
		seleciona, err := dbi.Database.Query(`select * from NOTIFICACAO order by data desc limit 1`)
		seleciona.Scan(&note.ID, &note.Data, &note.Lista)
		seleciona, err = dbi.Database.Query(`select * from EMAILENVIADO where idnotificacao = ` + strconv.Itoa(note.ID))
		
	}
	
	return err
}
