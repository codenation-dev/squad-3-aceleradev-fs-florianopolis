package main

import (
	"encoding/json"
	"io/ioutil"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	funcpublico "squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
	mail "squad-3-aceleradev-fs-florianopolis/services/MailSender/src"
)

/*O Banco Uati gostaria de monitorar de forma contínua e automatizada caso um de seus clientes
vire um funcionário público do estado de SP (http://www.transparencia.sp.gov.br/busca-agentes.html)
ou seja um bom cliente com um salário maior que 20 mil reais.*/

//Clientes do Banco Uati e que são funcionários públicos
func getTopClientsName() (nomes []string) {
	nomes, erro := funcpublico.GetClienteFuncPublico()
	if erro != nil {
		logs.Errorf("getNames", erro.Error())
	}
	//fmt.Println("Cliente funcionário público:", nomes)
	return nomes
}

//Funcionários públicos com maiores salários (acima de 20k, top 10)
func getTopIncomes() (nomes []string) {
	nomes, erro := funcpublico.GetTop10Incomes()
	if erro != nil {
		logs.Errorf("getNames", erro.Error())
	}
	//fmt.Println("Top 10 incomes: ", nomes)
	return nomes
}

func CreateJSONfile() {
	var request mail.Mailrequest
	request.Subject = "Take a look at UATI"
	request.Targets = []mail.Target{{"Roberta", "robertarl@gmail.com"},
		{"Luiz", "psychelipe@gmail.com"}, {"Rafael", "rfmf@protonmail.com"}, {"Rodrigo", "pp5ere@gmail.com"}}
	request.Message = "Esse email é um email de teste!"
	request.TopNames = getTopClientsName()
	request.Names = getTopIncomes()
	request.Link = "linkdobotao" //ALTERAR

	response, erro := json.Marshal(request)
	if erro != nil {
		logs.Errorf("createjsonfile", erro.Error())
	}
	erro = ioutil.WriteFile("mailrequest.json", response, 0644)
	if erro != nil {
		logs.Errorf("createjsonfile", erro.Error())
	}
}
