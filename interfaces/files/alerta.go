package main

import (
	"encoding/json"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	funcpublico "squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
	notificacao "squad-3-aceleradev-fs-florianopolis/interfaces/crud/notificacao"

	usuario "squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"

	mail "squad-3-aceleradev-fs-florianopolis/services/MailSender/src"
	"net/http"
	"bytes"
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

func sendMailRequest(r mail.Mailrequest) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	_,es := http.Post("http://127.0.0.1:8225/send", "application/json; charset=utf-8", b)
	if (es!=nil){
		logs.Errorf("MailConsumer - SendMailNotify",es.Error())
	}
}

func CreateNotify() {

	logs.Info("CreateJSONfile", "Generating JSON file for email service...")
	var request mail.Mailrequest
	request.Subject = "Novas oportunidades UATI"
	request.Targets = usuario.GetAllMails()
	request.Message = "Aqui temos as oportunidades da semana..."
	request.TopNames = getTopClientsName()
	request.Names = getTopIncomes()
	request.Link = "linkdobotao" //ALTERAR
	request.ID = notificacao.GetNextID()

	logs.Info("CreateNotify", "Updating notifications in DB...")
	notificacao.InsertNotificacao(request)
	logs.Info("CreateNotify", "Notifications up to date in DB!")

	sendMailRequest(request)

}
