package emailenviado

import (
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

