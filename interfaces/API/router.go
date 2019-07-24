package api

func (a *App) endpoints() {
	a.Router.HandleFunc("/auth",notImplemented)
	a.Router.HandleFunc("/warn",notImplemented)
	a.Router.HandleFunc("/warn/{id}",notImplemented)
	a.Router.HandleFunc("/mails",notImplemented)
	a.Router.HandleFunc("/mails/add",notImplemented)
	a.Router.HandleFunc("/mails/{id}/delete",notImplemented)
	a.Router.HandleFunc("/mails/{id}/update",notImplemented)
}

//Initialize the router
func Initialize(DBInterface DBI) App {

}
