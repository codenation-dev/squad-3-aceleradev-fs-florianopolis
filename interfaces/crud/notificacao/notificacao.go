package notificacao

import (
	"encoding/json"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	mail "squad-3-aceleradev-fs-florianopolis/services/MailSender/src"
	"strconv"
	"time"
)

//Insert Notificacao
func InsertNotificacao(request mail.Mailrequest) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	response, erro := json.Marshal(request)
	if erro != nil {
		logs.Errorf("InsertNotificacao", erro.Error())
	}
	_, erro = dbi.Database.Query(`INSERT INTO NOTIFICACAO (data, lista) VALUES(?, ?)`, time.Now().Format("2006-01-02 15:04:05"), response)
	if erro != nil {
		logs.Errorf("InsertNotificacao", erro.Error())
	}
	return erro
}

//Delete Notificacao by ID
func Delete(id int) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `DELETE FROM NOTIFICACAO WHERE id = ` + strconv.Itoa(id)
	erro = dbi.ExecQuery(squery)
	return erro
}
