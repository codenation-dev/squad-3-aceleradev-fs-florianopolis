package historico

import (
	"time"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"encoding/json"
)

//Insert historico
func Insert(hist *[]entity.FuncPublico) error {
	dbi, err := db.Init();if err != nil{
		return err
	}
	defer dbi.Database.Close()
	t := time.Now()
	formatTime := t.Format("2006-01-02 15:04:05")
	//formatTime := t.Format(time.RFC3339)
	json, err := json.Marshal(&hist);if err != nil{
		return err
	}
	 _, err = dbi.Database.Query(`INSERT INTO HISTORICO (data, json) VALUES(?, ?)`, formatTime, json)
	return err
}