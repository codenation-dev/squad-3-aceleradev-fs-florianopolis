package notificacao

import (
	"encoding/json"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
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

//Delete notificacao by ID
func Delete(id int) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `DELETE FROM NOTIFICACAO WHERE id = ` + strconv.Itoa(id)
	erro = dbi.ExecQuery(squery)
	return erro
}

//Get notificacao by ID
func Get(id int) (*entity.Notificacao, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	seleciona, erro := dbi.Database.Query(`SELECT * FROM notificacao WHERE id = ` + strconv.Itoa(id))
	var note entity.Notificacao
	if erro == nil {
		for seleciona.Next() {
			//erro := seleciona.Scan(&note.ID, &note.IDPessoa)
			if erro != nil {
				logs.Errorf("Get(notificação)", erro.Error())
			}
		}
	}

	return &note, erro
}
