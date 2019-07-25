package api

import (
	"github.com/gorilla/mux"
)

func (a *App) endpoints() {
	a.Router.HandleFunc("/auth",notImplemented)
	a.Router.HandleFunc("/warn",notImplemented)//data opcional vem no body da mensagem, se não preencher pega a última
	a.Router.HandleFunc("/warn/{idnotificacao}",notImplemented)
	a.Router.HandleFunc("/mails", a.mailGeneral)
	a.Router.HandleFunc("/mails/add",a.mailRegister)
	a.Router.HandleFunc("/mails/{id}/delete",a.mailDelete)
	a.Router.HandleFunc("/mails/{id}/update",a.mailEdit)
	a.Router.HandleFunc("/upload", a.uploadCSV)
}

//Initialize the router
func Initialize() *App {
	thisRouter := mux.NewRouter()
	thisApp := &App{Router: thisRouter}
	thisApp.initKey()
	thisApp.endpoints()
	return thisApp
}
