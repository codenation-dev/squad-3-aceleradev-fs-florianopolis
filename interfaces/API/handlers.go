package api

import("squad-3-aceleradev-fs-florianopolis/entities/logs"
entity "squad-3-aceleradev-fs-florianopolis/entities"
"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
"encoding/json"
"net/http"
"fmt"
"encoding/csv"
"io"
"io/ioutil"
"bytes"
"golang.org/x/crypto/bcrypt"
"github.com/gorilla/mux"
	"strconv"
	"strings"
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
			responseCodeResult(w, Error, E.Error())
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
	var UsuarioRequestUpdate entity.Usuario
	ids := mux.Vars(r)
	idStr := strings.Trim(ids["id"], " ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseCodeResult(w, Error, "Id não é um número")
	}else{
		UsuarioOnDataBase, err := usuario.GetUsuarioByID(id)
		if err != nil {
			responseCodeResult(w, Error, err.Error())
		}else{
			if UsuarioOnDataBase == nil{
				responseCodeResult(w, Empty, "Usuário não encontrado")
			}else{
				userJSON := json.NewDecoder(r.Body)
				err := userJSON.Decode(&UsuarioRequestUpdate)
				if err != nil {
					responseCodeResult(w, Error, err.Error())
				}else{
					UsuarioRequestUpdate.ID = id
					if UsuarioRequestUpdate.Validar(){
						err := usuario.Update(&UsuarioRequestUpdate)
						if err != nil {
							responseCodeResult(w, Error, err.Error())
						}else{
							responseCodeResult(w, Success, "Atualizado com Sucesso")
						}
					}
				}
			}
		}

	}
}

func (a *App) mailDelete(w http.ResponseWriter, r *http.Request)  {
	ids := mux.Vars(r)
	idStr := strings.Trim(ids["id"], " ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseCodeResult(w, Error, "Id não é um número")
	}else{
		UsuarioOnDataBase, err := usuario.GetUsuarioByID(id)
		if err != nil {
			responseCodeResult(w, Error, err.Error())
		}else{
			if UsuarioOnDataBase == nil{
				responseCodeResult(w, Empty, "Usuário não encontrado")
			}else{
				err := usuario.Delete(id)
				if err != nil {
					responseCodeResult(w, Error, err.Error())
				}else{
					responseCodeResult(w, Success, "Deletado com Sucesso")
				}
			}
		}

	}
}

func (a *App) mailRegister(w http.ResponseWriter, r *http.Request)  {
	
	decode := json.NewDecoder(r.Body)

	var Info MailType

	_ = decode.Decode(&Info)
	
	if (validateMailType(Info)){
		pass := generatePassword()
		crypted,_ := bcrypt.GenerateFromPassword([]byte(pass), 10)
		User := entity.Usuario{
			ID: 0,
		Usuario:Info.Name,
		Email: Info.Mail,
		Senha: string(crypted)}
		err := usuario.Insert(&User)
		
		if (err != nil){
			responseCodeResult(w, Error, err.Error())
		}
		
		u := passT{Subject: "Uati Suporte", 
		Target: MailType{Name:User.Usuario,
		Mail:User.Email},
		Message:pass}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(u)
		_,es := http.Post("http://127.0.0.1:8225/pass", "application/json; charset=utf-8", b)
		if (es!=nil){
			logs.Errorf("MailConsumer - SendMailWithPassError",es.Error())
			responseCodeResult(w, Error, "Cant Send Mail With Password")
		}
		responseCodeResult(w, Success, "Success")
	} else {
		responseCodeResult(w, Error, "Invalid Data")
	}

}

func (a *App) warnGeneral(w http.ResponseWriter, r *http.Request)  {
	var DataStruct DataEmailUsuario
	dateJSON := json.NewDecoder(r.Body)
	err := dateJSON.Decode(&DataStruct)
	if err != nil {
		responseCodeResult(w, Error, err.Error())
	}else{
		
	}
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
			responseCodeResult(w, Error, err.Error())
		}else{	
			err = ioutil.WriteFile("ClientList.json", j, 0644)
			if err != nil {
				responseCodeResult(w, Error, err.Error())
			}else{
				responseCodeResult(w, Success, "Arquivo salvo com sucesso")
			}
		}
	}else {
		responseCodeResult(w, Empty, "Nenhum dado encontrado")
	}
}

func responseCodeResult(w http.ResponseWriter, code int, msg string, tk ...string ){
	var response Result
	response.Code   = code
	response.Result = msg
	if (len(tk)>0) {
		response.Token = tk[0]
	}
	err := json.NewEncoder(w).Encode(response); if err != nil{
		logs.Errorf("App/Cant write respond", err.Error())	
	}
}
