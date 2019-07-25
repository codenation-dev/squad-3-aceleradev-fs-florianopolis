package api

import("squad-3-aceleradev-fs-florianopolis/entities/logs"
"squad-3-aceleradev-fs-florianopolis/entities"
"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
"encoding/json"
"net/http"
"fmt"
"encoding/csv"
"io"
"io/ioutil"
"bytes"
"golang.org/x/crypto/bcrypt"
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
		var response Result
		response.Result = "Nenhum usuário encontrado"
		err := json.NewEncoder(w).Encode(response); if err != nil{
			logs.Errorf("App/Cant write respond", err.Error())	
		}
	}
	
}

func (a *App) mailEdit(w http.ResponseWriter, r *http.Request)  {
	
}

func (a *App) mailDeleter(w http.ResponseWriter, r *http.Request)  {

}

func (a *App) mailRegister(w http.ResponseWriter, r *http.Request)  {
	
	decode := json.NewDecoder(r.Body)

	var Info MailType

	_ = decode.Decode(&Info)

	var Response Result
	
	if (validateMailType(Info)){
		pass := generatePassword()
		crypted,_ := bcrypt.GenerateFromPassword([]byte(pass), 2)
		User := entity.Usuario{
			ID: 0,
		Usuario:Info.Name,
		Email: Info.Mail,
		Senha: string(crypted)}
		usuario.Insert(&User)
		
		u := passT{Subject: "Uati Suporte", 
		Target: MailType{Name:User.Usuario,
		Mail:User.Email},
		Message:pass}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(u)
		rs,es := http.Post("http://127.0.0.1:8225/pass", "application/json; charset=utf-8", b)
		if (es!=nil){
			logs.Errorf("MailConsumer - SendMailWithPassError",es.Error())
		}
		cvt,_ := ioutil.ReadAll(rs.Body)
		logs.Info("MailConsumer - ResponseReceived", string(cvt))
		Response.Result = "Success"+pass 
	} else {
		Response.Result = "Fail"
	}

	 
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
	j, err := json.Marshal(structuredList)
	if err != nil {
		logs.Errorf("App/Cant encoding array to json", err.Error())
	}
	err = ioutil.WriteFile("ClientList.json", j, 0644)
	if err != nil {
		logs.Errorf("App/Cant write clientlist.json file on server", err.Error())
	}
}

func unauth(w http.ResponseWriter, r *http.Request) {

}

func internalError(w http.ResponseWriter, r *http.Request) {

}