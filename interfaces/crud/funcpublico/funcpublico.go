package funcpublico

import (
	"fmt"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
	"strings"
)

//Insert new funcionário público
func Insert(person *entity.FuncPublico) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	if erro != nil {
		logs.Errorf("Insert(FuncPublico)", erro.Error())
	}
	//NULLIF Prevents from creating empty string field in table FuncPublico
	squery := `INSERT INTO FUNCPUBLICO (nome, cargo, orgao, remuneracaodomes, ` +
		`redutorsalarial, totalliquido, updated, clientedobanco) VALUES(NULLIF("` +
		person.Nome + `",""), NULLIF("` + person.Cargo + `",""), NULLIF("` + person.Orgao + `",""), ` +
		strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) + `,` +
		strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) + `,` +
		strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) + `,` +
		strconv.FormatBool(person.Updated) + `,` +
		strconv.FormatBool(person.ClientedoBanco) + `);`
	result, erro := dbi.Database.Query(squery)
	defer result.Close()
	if erro != nil {
		logs.Errorf("Insert(FuncPublico)", erro.Error())
	}
	return erro
}

//Delete funcionário público by ID
func Delete(id int) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	result, erro := dbi.Database.Query(`DELETE FROM FUNCPUBLICO WHERE id = ` + strconv.Itoa(id))
	defer result.Close()
	return erro
}

//Update funcionário público
func Update(person *entity.FuncPublico) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	if erro != nil {
		logs.Errorf("Insert(FuncPublico)", erro.Error())
	}
	//NULLIF Prevents from creating empty string field in table FuncPublico
	squery := `UPDATE FUNCPUBLICO SET nome = NULLIF("` + person.Nome + `",""), cargo = NULLIF("` + person.Cargo +
		`",""), orgao = NULLIF("` + person.Orgao + `",""), remuneracaodomes = ` + strconv.FormatFloat(person.Remuneracaodomes, 'E', -1, 64) +
		`, redutorsalarial = ` + strconv.FormatFloat(person.RedutorSalarial, 'E', -1, 64) +
		`, totalliquido = ` + strconv.FormatFloat(person.TotalLiquido, 'E', -1, 64) +
		`, updated = ` + strconv.FormatBool(person.Updated) +
		`, clientedobanco = ` + strconv.FormatBool(person.ClientedoBanco) +
		` WHERE id = ` + strconv.Itoa(person.ID)
	result, erro := dbi.Database.Query(squery)
	defer result.Close()
	if erro != nil {
		logs.Errorf("Insert(FuncPublico)", erro.Error())
	}
	return erro
}

//UpdateAllSetRemuneracaodoMes atualiza os valores dos que não são mais funcionários públicos para 0
func UpdateAllSetRemuneracaodoMes(remuneracaodomes float64) error { //CHANGED from totalliquido to remuneracaodomes
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `UPDATE FUNCPUBLICO SET remuneracaodomes = ` + strconv.FormatFloat(remuneracaodomes, 'E', -1, 64) +
		`WHERE updated = ` + strconv.FormatBool(false) //CHANGED from totalliquido to remuneracaodomes
	result, erro := dbi.Database.Query(squery)
	defer result.Close()
	return erro
}

//UpdateAllSetFlagUpdated seta a flag updated de todos conforme valor passado
func UpdateAllSetFlagUpdated(flag bool) error {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `UPDATE FUNCPUBLICO SET updated = ` + strconv.FormatBool(flag)
	result, erro := dbi.Database.Query(squery)
	defer result.Close()
	return erro
}

//Get funcionário público by ID

func Get(id int) (*entity.FuncPublico, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `SELECT * FROM FUNCPUBLICO WHERE id = ` + strconv.Itoa(id)
	seleciona, erro := dbi.Database.Query(squery)
	var person entity.FuncPublico
	defer seleciona.Close()
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				logs.Errorf("Get(funcpublico)", erro.Error())
			}
		}
	}
	return &person, erro
}

//GetByName funcionário público by name
func GetByName(name string) (*entity.FuncPublico, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	name = strings.Replace(name, "'", "''", 1) //prevent from single quotes in names (Escape character)
	squery := `SELECT * FROM FUNCPUBLICO WHERE nome = "` + strings.Trim(name, " ") + `"`
	fmt.Println(squery)
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var person entity.FuncPublico
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				logs.Errorf("GetByName(funcpublico)", erro.Error())
			}
		}
	}
	fmt.Println(person)
	return &person, erro
}

//GetAllFuncPublico get all funcionário público
func GetAllFuncPublico() (*[]entity.FuncPublico, error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `SELECT * FROM FUNCPUBLICO WHERE remuneracaodomes > 0 ORDER BY totalliquido DESC` //CHANGED from totalliquido to remuneracaodomes
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var person entity.FuncPublico
	var persons []entity.FuncPublico
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			persons = append(persons, person)
			if erro != nil {
				logs.Errorf("GetAllFuncPublico", erro.Error())
				break
			}
		}
	}
	return &persons, erro
}

//GetClienteFuncPublico cliente que é funcionário público
func GetClienteFuncPublico() (nomes []string, erro error) {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	squery := `SELECT * FROM FUNCPUBLICO WHERE clientedobanco = TRUE LIMIT 10`
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var person entity.FuncPublico
	var names []string
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.Nome, &person.Cargo, &person.Orgao,
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Updated, &person.ClientedoBanco)
			if erro != nil {
				logs.Errorf("GetClienteFuncPublico", erro.Error())
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
	squery := `SELECT nome FROM FUNCPUBLICO ORDER BY totalliquido DESC LIMIT 10`
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var names []string
	var name string
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&name)
			if erro != nil {
				logs.Errorf("GetTop10Incomes", erro.Error())
			}
			names = append(names, name)
		}
	}
	return names, erro
}
