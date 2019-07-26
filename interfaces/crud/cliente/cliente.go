package cliente

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
)


func Insert(client *entity.Cliente) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Insert(Cliente)", erro.Error())
	}
	defer dbi.Database.Close()
	logs.Info("Insert(Cliente)", "Inserting new Cliente in DB...")
	_, erro = dbi.Database.Query(`INSERT INTO CLIENTE (nome, idfuncpublico) VALUES(?, ?)`, client.Nome, client.IDFuncPublico)
	return erro
}