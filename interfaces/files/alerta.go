package main

import (
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	notificacao "squad-3-aceleradev-fs-florianopolis/interfaces/crud/notificacao"

	usuario "squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	utils "squad-3-aceleradev-fs-florianopolis/utils"

)

/*O Banco Uati gostaria de monitorar de forma contínua e automatizada caso um de seus clientes
vire um funcionário público do estado de SP (http://www.transparencia.sp.gov.br/busca-agentes.html)
ou seja um bom cliente com um salário maior que 20 mil reais.*/

//Clientes do Banco Uati e que são funcionários públicos

func CreateNotify() {

	logs.Info("CreateJSONfile", "Generating JSON file for email service...")
	request := utils.RequestCreator(usuario.GetAllMails(),notificacao.GetNextID())

	logs.Info("CreateNotify", "Updating notifications in DB...")
	notificacao.InsertNotificacao(request)
	logs.Info("CreateNotify", "Notifications up to date in DB!")

	utils.SendMailRequest(request)
}

