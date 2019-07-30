package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/emailenviado"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/notificacao"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	utils "squad-3-aceleradev-fs-florianopolis/utils"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotImplemented But Success")
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {

	decode := json.NewDecoder(r.Body)

	var cred credentials

	_ = decode.Decode(&cred)

	var Response Result

	if a.tryLogin(&cred) {
		T, E := a.GenerateJWT(cred.Usermail)
		if E == nil {
			Response.Result = "Login Success"
			Response.Token = T
			Response.Code = Success
		} else {
			logs.Errorf("App/Cant create JWT Token", E.Error())
			responseCodeResult(w, Error, E.Error())
		}
	} else {
		responseCodeResult(w, Error, "Invalid Credentials")
	}
	json.NewEncoder(w).Encode(Response)
}

func (a *App) mailGeneral(w http.ResponseWriter, r *http.Request) {
	usersEmails := usuario.GetAllMails()
	if usersEmails != nil {
		var Response Result
		Response.Usermails = &usersEmails
		Response.Code = Success
		Response.Result = "Success"
		Response.Token = a.GetToken(context.Get(r, "token").(*jwt.Token))
		err := json.NewEncoder(w).Encode(Response)
		if err != nil {
			logs.Errorf("App/Cant write respond", err.Error())
		}
	} else {
		responseCodeResult(w, Empty, "Nenhum dado encontrado", a.GetToken(context.Get(r, "token").(*jwt.Token)))
	}

}

func (a *App) mailEdit(w http.ResponseWriter, r *http.Request) {
	var UsuarioRequestUpdate entity.Usuario
	ids := mux.Vars(r)
	idStr := strings.Trim(ids["id"], " ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseCodeResult(w, 5, "Id não é um número", a.GetToken(context.Get(r, "token").(*jwt.Token)))
	} else {
		UsuarioOnDataBase, err := usuario.GetUsuarioByID(id)
		if err != nil {
			responseCodeResult(w, 6, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
		} else {
			if UsuarioOnDataBase == nil {
				responseCodeResult(w, Empty, "Usuário não encontrado")
			} else {
				userJSON := json.NewDecoder(r.Body)
				err := userJSON.Decode(&UsuarioRequestUpdate)
				if err != nil {
					responseCodeResult(w, 7, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
				} else {
					UsuarioRequestUpdate.ID = id
					Temp := UsuarioRequestUpdate.Senha
					if Temp == "" {
						err = usuario.UpdateWithoutPass(&UsuarioRequestUpdate)
					} else {

						pass, err := bcrypt.GenerateFromPassword([]byte(Temp), 10)
						if err != nil {
							responseCodeResult(w, 8, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
						}
						UsuarioRequestUpdate.Senha = string(pass)
						dbi, erro := db.Init() //changed to mock DB for unit test
						if erro != nil {       //changed to mock DB for unit test
							logs.Errorf("mailEdit(handlers)", erro.Error()) //changed to mock DB for unit test
						} //changed to mock DB for unit test
						defer dbi.Database.Close()                       //changed to mock DB for unit test
						err = usuario.Update(&UsuarioRequestUpdate, dbi) //changed to mock DB for unit test
					}

					if err != nil {
						responseCodeResult(w, 9, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
					} else {
						responseCodeResult(w, Success, "Atualizado com Sucesso", a.GetToken(context.Get(r, "token").(*jwt.Token)))
					}
				}
			}
		}

	}
}

func (a *App) mailDelete(w http.ResponseWriter, r *http.Request) {
	ids := mux.Vars(r)
	idStr := strings.Trim(ids["id"], " ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responseCodeResult(w, Error, "Id não é um número", a.GetToken(context.Get(r, "token").(*jwt.Token)))
	} else {
		UsuarioOnDataBase, err := usuario.GetUsuarioByID(id)
		if err != nil {
			responseCodeResult(w, Error, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
		} else {
			if UsuarioOnDataBase == nil {
				responseCodeResult(w, Empty, "Usuário não encontrado", a.GetToken(context.Get(r, "token").(*jwt.Token)))
			} else {
				dbi, erro := db.Init() //changed to mock DB for unit test
				if erro != nil {       //changed to mock DB for unit test
					logs.Errorf("mailDelete(handlers)", erro.Error()) //changed to mock DB for unit test
				} //changed to mock DB for unit test
				defer dbi.Database.Close()     //changed to mock DB for unit test
				err := usuario.Delete(id, dbi) //changed to mock DB for unit test
				if err != nil {
					responseCodeResult(w, Error, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
				} else {
					responseCodeResult(w, Success, "Deletado com Sucesso", a.GetToken(context.Get(r, "token").(*jwt.Token)))
				}
			}
		}

	}
}

func (a *App) mailRegister(w http.ResponseWriter, r *http.Request) {

	decode := json.NewDecoder(r.Body)

	var Info entity.Target

	_ = decode.Decode(&Info)

	if validateMailType(Info) {
		pass := utils.GeneratePassword()
		crypted, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
		User := entity.Usuario{
			ID:      0,
			Usuario: Info.Name,
			Email:   Info.Mail,
			Senha:   string(crypted)}
		dbi, erro := db.Init()     //changed to mock DB for unit test
		defer dbi.Database.Close() //changed to mock DB for unit test
		if erro != nil {           //changed to mock DB for unit test
			logs.Errorf("mailRegister(handlers)", erro.Error()) //changed to mock DB for unit test
		} //changed to mock DB for unit test
		err := usuario.Insert(&User, dbi) //changed to mock DB for unit test

		if err != nil {
			responseCodeResult(w, Error, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
		}
		tgt := entity.Target{Name: User.Usuario,
			Mail: User.Email}
		u := passT{Subject: "Uati Suporte",
			Target:  tgt,
			Message: pass}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(u)
		_, es := http.Post("http://127.0.0.1:8225/pass", "Application/json; charset=utf-8", b)
		if es != nil {
			logs.Errorf("MailConsumer - SendMailWithPassError", es.Error())
			responseCodeResult(w, Error, "Cant Send Mail With Password", a.GetToken(context.Get(r, "token").(*jwt.Token)))
		} else {
			responseCodeResult(w, Success, "Success", a.GetToken(context.Get(r, "token").(*jwt.Token)))
			request := utils.RequestCreator([]entity.Target{tgt}, notificacao.GetLastID())
			utils.SendMailRequest(request)
		}
	} else {
		responseCodeResult(w, Error, "Invalid Data", a.GetToken(context.Get(r, "token").(*jwt.Token)))
	}

}

func (a *App) warnGeneral(w http.ResponseWriter, r *http.Request) {
	var DataStruct DataEmailUsuario
	dateJSON := json.NewDecoder(r.Body)
	err := dateJSON.Decode(&DataStruct)
	if err != nil && err != io.EOF {
		responseCodeResult(w, Error, err.Error())
	} else {
		var warn Warn
		Notificacao, err := notificacao.Get(DataStruct.Data)
		if err != nil {
			responseCodeResult(w, Error, err.Error())
		}
		Emails := emailenviado.GetAll(Notificacao.ID)
		warn.ID = Notificacao.ID
		warn.Lista = Notificacao.Lista
		warn.Data = Notificacao.Data
		warn.EmailsEnviados = Emails
		var Response Result
		Response.Warn = &warn
		Response.Code = Success
		Response.Result = "Sucesso"
		Response.Token = a.GetToken(context.Get(r, "token").(*jwt.Token))
		err = json.NewEncoder(w).Encode(Response)
		if err != nil {
			logs.Errorf("App/Cant write respond", err.Error())
		}
	}
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
			logs.Errorf("App/Cant read file", err.Error())
		}
		cliente := ListaClientes{
			Nome: line[0],
		}
		structuredList = append(structuredList, cliente)
	}
	if structuredList != nil {
		j, err := json.Marshal(structuredList)
		if err != nil {
			responseCodeResult(w, Error, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
		} else {
			err = ioutil.WriteFile("ClientList.json", j, 0644)
			if err != nil {
				responseCodeResult(w, Error, err.Error(), a.GetToken(context.Get(r, "token").(*jwt.Token)))
			} else {
				responseCodeResult(w, Success, "Arquivo salvo com sucesso", a.GetToken(context.Get(r, "token").(*jwt.Token)))
			}
		}
	} else {
		responseCodeResult(w, Empty, "Nenhum dado encontrado", a.GetToken(context.Get(r, "token").(*jwt.Token)))
	}
}

func responseCodeResult(w http.ResponseWriter, code int, msg string, tk ...string) {
	var response Result
	response.Code = code
	response.Result = msg
	if len(tk) > 0 {
		response.Token = tk[0]
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logs.Errorf("App/Cant write respond", err.Error())
	}
}
