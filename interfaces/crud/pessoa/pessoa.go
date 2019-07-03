package pessoa

import (
	db "squad-3-aceleradev-fs-florianopolis/db"
	"squad-3-aceleradev-fs-florianopolis/entity"
	"strconv"
)

//Insert new pessoa
func Insert(person *entity.Pessoa) error {
	erro := db.ExecutaComando("insert into pessoa ( nome, cargo, orgao, remuneracaodomes, 
		redutorsalarial, totalliquido, update, clientedobanco) values(" 
		+ person.Nome + "', '" + person.Cargo + "','" + person.Orgao + "'," + strconv.FormatFloat(person.Remuneracaodomes) + "'," 
		+ strconv.FormatFloat(person.RedutorSalarial) + "'," + strconv.FormatFloat(person.TotalLiquido) + "',"
		+ strconv.FormatBool(person.Update) + "',"+ strconv.FormatBool(person.ClientedoBanco) + "'," + ");")
	return erro

//Delete pessoa by ID
func Delete(id int) error {
	erro := db.ExecutaComando("delete from pessoa where id = " + strconv.Itoa(id))
	return erro
}

//Update pessoa
func Update(person *entity.Pessoa) error {
	erro := db.ExecutaComando("update pessoa set idarquivotransparencia = " + strconv.Itoa(person.IDArquivoTransparencia) 
	+ ", nome = '" + person.Nome + "', cargo = '" + person.Cargo + "', orgao = '" + person.Orgao + "', remuneracaodomes = " 
	+ strconv.FormatFloat(person.Remuneracaodomes) + "', redutorsalarial = " + strconv.FormatFloat(person.RedutorSalarial) 
	+ "', totalliquido = " + strconv.FormatFloat(person.TotalLiquido) + "', update = " + strconv.FormatBool(person.Update) 
	+ "', clientedobanco = " + strconv.FormatBool(person.ClientedoBanco) + " where id = " + strconv.Itoa(person.ID))
	return erro
}

//Get pessoa by ID
func Get(id int) (*entity.Pessoa, error) {
	db, erro := db.Conect()
	seleciona, erro := db.Query("select * from pessoa where id = " + strconv.Itoa(id))
	var person entity.Pessoa
	if erro == nil {

		for seleciona.Next() {
			erro := seleciona.Scan(&person.ID, &person.IDArquivoTransparencia, &person.Nome, &person.Cargo, &person.Orgao, 
				&person.Remuneracaodomes, &person.RedutorSalarial, &person.TotalLiquido, &person.Update, &person.ClientedoBanco)
			if erro != nil {
				panic(erro.Error())
			}
		}
	}
	defer db.Close()
	return &person, erro
}

