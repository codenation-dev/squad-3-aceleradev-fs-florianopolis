// API
//
// This aplication compares a given list of clients to data taken from the open
// API about public servers of the state of São Paulo, Brazil, and serves
// alerts and analytics to the users.
//
//	Schemes: http
//	Host: localhost:8921
//	Version: 0.0.1
//	License: MIT http://opensource.org/liceses/MIT
//	Contact: Rafael<rfmf@protonmail.com>
//
//	Consumes:
//	- text/plain
//
//	Produces:
//	- text/plain
//
//	swagger:meta

package api

import (
	"github.com/gorilla/mux"
)

func (a *App) endpoints() {

	a.Router.HandleFunc("/auth", a.restricted(a.login))
	a.Router.HandleFunc("/warn", a.restricted(a.warnGeneral)) //data em um json opcional vem no body da mensagem, se não preencher pega a última
	a.Router.HandleFunc("/mails", a.restricted(a.mailGeneral))
	a.Router.HandleFunc("/mails/add", a.restricted(a.mailRegister))
	a.Router.HandleFunc("/mails/{id}/delete", a.restricted(a.mailDelete))
	a.Router.HandleFunc("/mails/{id}/update", a.restricted(a.mailEdit))
	a.Router.HandleFunc("/upload", a.restricted(a.uploadCSV))
	a.Router.HandleFunc("/tables", a.restricted(a.serveDSTables))
}

//Initialize the router
func Initialize() *App {
	thisRouter := mux.NewRouter()
	thisApp := &App{Router: thisRouter}
	thisApp.initKey()
	thisApp.endpoints()
	return thisApp
}
