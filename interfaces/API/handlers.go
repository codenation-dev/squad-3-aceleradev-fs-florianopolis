package api

import(
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"encoding/json"
	"net/http"
	"fmt"
	"encoding/csv"
	"io"
	"io/ioutil"
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
	usersEmails := usuario.GetAllMails()
	if usersEmails != nil{
		err := json.NewEncoder(w).Encode(usersEmails); if err != nil{
			logs.Errorf("App/Cant write respond", err.Error())	
		}
	}else {
		responseCodeResult(w, Empty, "Nenhum dado encontrado")
	}
	
}

func (a *App) mailEdit(w http.ResponseWriter, r *http.Request)  {
	var Usuario entity.Usuario
	userJSON := json.NewDecoder(r.Body)
	err := userJSON.Decode(&Usuario);if err != nil {
		responseCodeResult(w, Error, err.Error())
	}

	if Usuario.Validar(){
		usuario.Update(&Usuario)	
	}

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
	var structuredList []ListaClientes
	for {
		line, err := list.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Errorf("App/Cant read file",err.Error())
		}
		cliente := ListaClientes{
			Nome:    line[0],
		}
		structuredList = append(structuredList, cliente)
	}
	if structuredList != nil{
		j, err := json.Marshal(structuredList)
		if err != nil {
			logs.Errorf("App/Cant encoding array to json", err.Error())
		}
		err = ioutil.WriteFile("ClientList.json", j, 0644)
		if err != nil {
			logs.Errorf("App/Cant write clientlist.json file on server", err.Error())
		}
	}else {
		responseCodeResult(w, Empty, "Nenhum dado encontrado")
	}
}

func unauth(w http.ResponseWriter, r *http.Request) {

}

func internalError(w http.ResponseWriter, r *http.Request) {

}

func responseCodeResult(w http.ResponseWriter, code int, msg string ){
	var response Result
	response.Code   = code
	response.Result = msg
	err := json.NewEncoder(w).Encode(response); if err != nil{
		logs.Errorf("App/Cant write respond", err.Error())	
	}
}