package api

import(
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"encoding/json"
	"net/http"
	"fmt"
)

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"NotImplemented But Success")
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {

	decode := json.NewDecoder(r.Body)

	var cred credentials

	_ = decode.Decode(&cred)
	
	var Response Result

	if (a.tryLogin(&cred)){
		T,E := a.GenerateJWT(cred.Usermail)
		if (E==nil){
			Response.Result = "Login Success"
			Response.Token = T
		} else {
			logs.Errorf("App/Cant create JWT Token",E.Error())
			Response.Result = "Login Fail! Internal Error"
		}	
	} else {
		Response.Result = "Login Fail! Invalid Credentials"
	}
	json.NewEncoder(w).Encode(Response)
}

func (a *App) mailGeneral(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailEdit(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailDeleter(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailRegister(w http.ResponseWriter, r *http.Request)  {
	//TruePass := bcrypt.GenerateFromPassword([]byte(), 2)
}

func (a *App) warnGeneral(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) warnDetail(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) uploadCSV(w http.ResponseWriter, r *http.Request) {
	list := csv.NewReader(r.Body)
	var structuredList []listaClientes
	for {
		line, err := list.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Errorf("App/Cant read file",err.Error())
		}
		cliente := listaClientes{
			Nome:    line[0],
		}
		structuredList = append(structuredList, cliente)
	}
	j, err := json.Marshal(structuredList)
	if err != nil {
		logs.Errorf("App/Cant encoding array to json", err.Error())
	}
	err = ioutil.WriteFile("clientlist.json", j, 0644)
	if err != nil {
		logs.Errorf("App/Cant write clientlist.json file on server", err.Error())
	}
}

func unauth(w http.ResponseWriter, r *http.Request) {

}

func internalError(w http.ResponseWriter, r *http.Request) {

}