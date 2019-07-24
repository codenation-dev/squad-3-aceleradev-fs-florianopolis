package api

import("squad-3-aceleradev-fs-florianopolis/entities/logs"
"encoding/json"
"net/http"
"fmt")

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"NotImplemented But Success")
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {

decode := json.NewDecoder(r.Body)

var cred credentials

_ = decode.Decode(&cred)

if(a.loginAttempt(&cred)){
		T,E := a.GenerateJWT(cred.Usermail)
		if (E!=nil){
			logs.Errorf("App/Cant create JWT Token",E.Error())
		}
		fmt.Fprintf(w,fmt.Sprintf("this is your token: %s",string(T)))
} else {
	fmt.Fprintf(w,"Login Error")
}

}

func (a *App) mailGeneral(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailEdit(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailDeleter(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailRegister(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) warnGeneral(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) warnDetail(w http.ResponseWriter, r *http.Request)  {

}

func unauth(w http.ResponseWriter, r *http.Request) {

}

func internalError(w http.ResponseWriter, r *http.Request) {

}