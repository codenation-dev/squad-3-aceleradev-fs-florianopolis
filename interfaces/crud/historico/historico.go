package historico

import (
	"time"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
	"encoding/json"
)

func Insert(hist *[]entity.Historico) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	t := time.Now()
	formatTime := t.Format(time.RFC1123)
	json, err := json.Marshal(hist)
	squery := "INSERT INTO HISTORICO (data, json) VALUES('" + formatTime +
	 ");"
	erro = dbi.ExecQuery(squery)
	return erro
}