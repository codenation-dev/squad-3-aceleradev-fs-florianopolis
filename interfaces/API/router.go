package api

import ("github.com/gorilla/mux")

func (a *App) endpoints() {
	a.Router.HandleFunc("/auth", a.login)
	a.Router.HandleFunc("/warn",notImplemented)
	a.Router.HandleFunc("/warn/{id}",notImplemented)
	a.Router.HandleFunc("/mails",notImplemented)
	a.Router.HandleFunc("/mails/add",notImplemented)
	a.Router.HandleFunc("/mails/{id}/delete",notImplemented)
	a.Router.HandleFunc("/mails/{id}/update",notImplemented)
	a.Router.HandleFunc("/upload",notImplemented)
}

//Initialize the router
func Initialize() *App {
	thisRouter := mux.NewRouter()
	thisApp := &App{Router: thisRouter}
	thisApp.initKey()
	thisApp.endpoints()
	return thisApp
}
