package funcpublico

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
)

//Insert new funcionário público
//Alterado para receber conexão como parâmetro, para não ficar dando Db.Init() em cada chamada
func Insert(person *entity.FuncPublico, dbi *db.MySQLDatabase) error {
	//dbi, erro := db.Init()
	squery := "INSERT INTO FUNCPUBLICO (nome, cargo, orgao, remuneracaodomes, " +
		"redutorsalarial, totalliquido, updated, clientedobanco) VALUES('" +
		person.Nome + "', '" + person.Cargo + "','" + person.Orgao + "'," + strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) + " ," +
		strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) + "," + strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) + "," +
		strconv.FormatBool(person.Updated) + "," + strconv.FormatBool(person.ClientedoBanco) + ");"
	erro, _ := dbi.ExecQuery(squery)
	return erro
}

//Delete funcionário público by ID
//Alterado para receber conexão como parâmetro, para não ficar dando Db.Init() em cada chamada
func Delete(id int, dbi *db.MySQLDatabase) error {
	//dbi, erro := db.Init()
	erro, _ := dbi.ExecQuery("DELETE FROM FUNCPUBLICO WHERE id = " + strconv.Itoa(id))
	//defer dbi.Database.Close()
	return erro
}

//Update funcionário público
//Alterado para receber conexão como parâmetro, para não ficar dando Db.Init() em cada chamada
func Update(person *entity.FuncPublico, dbi *db.MySQLDatabase) error {
	//dbi, erro := db.Init()
	squery := "UPDATE FUNCPUBLICO SET nome = '" + person.Nome + "', cargo = '" + person.Cargo +
		"', orgao = '" + person.Orgao + "', remuneracaodomes = " + strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) +
		", redutorsalarial = " + strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) +
		", totalliquido = " + strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) + ", updated = " + strconv.FormatBool(person.Updated) +
		", clientedobanco = " + strconv.FormatBool(person.ClientedoBanco) + " WHERE id = " + strconv.Itoa(person.ID)
	erro, _ := dbi.ExecQuery(squery)
	//defer dbi.Database.Close()
	return erro
}

//Get funcionário público by ID
//Alterado para receber conexão como parâmetro, para não ficar dando Db.Init() em cada chamada
func Get(id int, dbi *db.MySQLDatabase) (*entity.FuncPublico, error) {
	//dbi, erro := db.Init()
	squery := "SELECT * FROM FUNCPUBLICO WHERE id = " + strconv.Itoa(id)
	erro, seleciona := dbi.ExecQuery(squery)
	var person entity.FuncPublico
	//defer dbi.Database.Close()
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				panic(erro.Error())
			}
		}
	}
	return &person, erro
}

//Get funcionário público by name
//Alterado para receber conexão como parâmetro, para não ficar dando Db.Init() em cada chamada
func GetByName(name string, dbi *db.MySQLDatabase) (*entity.FuncPublico, error) {
	//dbi, erro := db.Init()
	squery := "SELECT * FROM FUNCPUBLICO WHERE nome = '" + name + "'"
	erro, seleciona := dbi.ExecQuery(squery)
	//defer dbi.Database.Close()
	var person entity.FuncPublico
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				panic(erro.Error())
			}
		}
	}
	return &person, erro
}
