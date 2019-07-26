package cliente

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strings"
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

func GetByName(name string) (*entity.Cliente, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	name = strings.Replace(name, "'", "''", 1) //prevent from single quotes in names (Escape character)
	squery := `SELECT * FROM CLIENTE WHERE nome = "` + strings.Trim(name, " ") + `"`
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var cliente entity.Cliente
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&cliente.ID, &cliente.Nome, &cliente.IDFuncPublico)
			if erro != nil {
				logs.Errorf("GetByName(funcpublico)", erro.Error())
			}
		}
	}
	return &cliente, erro
}
