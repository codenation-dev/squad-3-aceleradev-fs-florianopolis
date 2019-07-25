package main

import (
	"encoding/json"
	"io/ioutil"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	funcpublico "squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
	notificacao "squad-3-aceleradev-fs-florianopolis/interfaces/crud/notificacao"
	usuario "squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
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
	return nomes
}

//Funcionários públicos com maiores salários (acima de 20k, top 10)
func getTopIncomes() (nomes []string) {
	nomes, erro := funcpublico.GetTop10Incomes()
	if erro != nil {
		logs.Errorf("getNames", erro.Error())
	}
	return nomes
}

func CreateJSONfile() {
	logs.Info("CreateJSONfile", "Generating JSON file for email service...")
	var request mail.Mailrequest
	request.Subject = "Novas oportunidades UATI"
	request.Targets = usuario.GetAllMails()
	request.Message = "Aqui temos as oportunidades da semana..."
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
	logs.Info("CreateJSONfile", "Updating notifications in DB...")
	notificacao.InsertNotificacao(request)
	logs.Info("CreateJSONfile", "Notifications up to date in DB!")
}
