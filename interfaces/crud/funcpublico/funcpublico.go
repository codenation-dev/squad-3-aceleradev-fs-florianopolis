package funcpublico

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
)

//Insert new funcionário público
//func Insert(person *entity.FuncPublico, dbi *db.MySQLDatabase) error {
func Insert(person *entity.FuncPublico) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "INSERT INTO FUNCPUBLICO (nome, cargo, orgao, remuneracaodomes, " +
		"redutorsalarial, totalliquido, updated, clientedobanco) VALUES('" +
		person.Nome + "', '" + person.Cargo + "','" + person.Orgao + "'," + strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) + " ," +
		strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) + "," + strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) + "," +
		strconv.FormatBool(person.Updated) + "," + strconv.FormatBool(person.ClientedoBanco) + ");"
	erro = dbi.ExecQuery(squery)
	return erro
}

//Delete funcionário público by ID
//func Delete(id int, dbi *db.MySQLDatabase) error {
func Delete(id int) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	erro = dbi.ExecQuery("DELETE FROM FUNCPUBLICO WHERE id = " + strconv.Itoa(id))
	return erro
}

//Update funcionário público
//func Update(person *entity.FuncPublico, dbi *db.MySQLDatabase) error {
func Update(person *entity.FuncPublico) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "UPDATE FUNCPUBLICO SET nome = '" + person.Nome + "', cargo = '" + person.Cargo +
		"', orgao = '" + person.Orgao + "', remuneracaodomes = " + strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) +
		", redutorsalarial = " + strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) +
		", totalliquido = " + strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) + ", updated = " + strconv.FormatBool(person.Updated) +
		", clientedobanco = " + strconv.FormatBool(person.ClientedoBanco) + " WHERE id = " + strconv.Itoa(person.ID)
	erro = dbi.ExecQuery(squery)
	return erro
}

//UpdateAllSetTotalLiquido atualiza os valores dos que não são mais funcionários públicos para 0
func UpdateAllSetTotalLiquido(totalLiquido float64) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "UPDATE FUNCPUBLICO SET totalliquido = " + strconv.FormatFloat(totalLiquido, 'E', -1, 64) +
		"WHERE updated = " + strconv.FormatBool(false)
	erro = dbi.ExecQuery(squery)
	return erro
}

//UpdateAllSetFlagUpdated seta a flag updated de todos conforme valor passado
func UpdateAllSetFlagUpdated(flag bool) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "UPDATE FUNCPUBLICO SET updated = " + strconv.FormatBool(flag)
	erro = dbi.ExecQuery(squery)
	return erro
}

//Get funcionário público by ID
//func Get(id int, dbi *db.MySQLDatabase) (*entity.FuncPublico, error) {
func Get(id int) (*entity.FuncPublico, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "SELECT * FROM FUNCPUBLICO WHERE id = " + strconv.Itoa(id)
	seleciona, erro := dbi.Database.Query(squery)
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

//Get funcionário público by name
func GetByName(name string) (*entity.FuncPublico, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "SELECT * FROM FUNCPUBLICO WHERE nome = '" + name + "'"
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
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

//Get cliente que é funcionário público
func GetClienteFuncPublico() (nomes []string, erro error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "SELECT * FROM FUNCPUBLICO WHERE clientedobanco = TRUE"
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var person entity.FuncPublico
	var names []string
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				panic(erro.Error())
			}
			names = append(names, person.Nome)
		}
	}
	return names, erro
}

//Get top10 func publico based on income
func GetTop10Incomes() (nomes []string, erro error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := "SELECT nome FROM FUNCPUBLICO ORDER BY totalliquido DESC LIMIT 10"
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var names []string
	var name string
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&name)
			if erro != nil {
				panic(erro.Error())
			}
			names = append(names, name)
		}
	}
	return names, erro
}
