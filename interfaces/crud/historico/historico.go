package historico

import (
	"encoding/json"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"time"
)

//Insert historico da tabela Funcpublico
func Insert(hist *[]entity.FuncPublico) error {
	dbi, err := db.Init()
	if err != nil {
		return err
	}
	defer dbi.Database.Close()
	t := time.Now()
	formatTime := t.Format("2006-01-02 15:04:05")
	json, err := json.Marshal(&hist)
	if err != nil {
		return err
	}
	result, err := dbi.Database.Query(`INSERT INTO HISTORICO (data, json) VALUES(?, ?)`, formatTime, json)
	defer result.Close()
	return err
}
