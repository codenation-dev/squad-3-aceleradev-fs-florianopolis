package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/emailenviado"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/notificacao"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	utils "squad-3-aceleradev-fs-florianopolis/utils"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// swagger:operation POST /auth
//
// Checks the user credentials and returns a token.
//
// consumes:
// - text/plain
// produces:
// - text/plain
// parameter:
// - Username: user name
// - Password: user password
// responses:
//  '200':
//		Result: result of request.
//		Code: code of response.
//		Token: The session token.
//  '401':
//		Result: Unauthorized Access.
//		Code: 9
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

// swagger:operation POST /mails
//
// Returns all e-mails registered.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response.
//   token: session token.
//   UsermailList: list of registered e-mails.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
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

// swagger:operation POST /mails/{id}/update id
//
// Updates user data.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response.
//   token: session token.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
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

// swagger:operation POST /mails/{id}/delete id
//
// Delete user data.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response
//   token: session token.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
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
				//err := usuario.Delete(id)
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

// swagger:operation POST /mails/add
//
// Register user.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response
//   token: session token.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
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

// swagger:operation POST /warn
//
// Provides warnings to the user.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response
//   token: session token.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
func (a *App) warnGeneral(w http.ResponseWriter, r *http.Request) {
	var DataStruct DataEmailUsuario
	var Data time.Time
	dateJSON := json.NewDecoder(r.Body)
	err := dateJSON.Decode(&DataStruct)
	if err != nil && err != io.EOF {
		responseCodeResult(w, Error, err.Error())
	} else {
		var warn Warn
		var Notificacao *entity.Notificacao
		if DataStruct.Data != Data {
			Notificacao, err = notificacao.Get(DataStruct.Data)
		} else {
			if DataStruct.ID != 0 {
				Notificacao, err = notificacao.GetByID(DataStruct.ID)
			} else {
				Notificacao, err = notificacao.Get(DataStruct.Data)
			}
		}
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

// swagger:operation POST /upload
//
// Allows the user to upload a client list to the server.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   Result: result of request.
//   Code: code of response
//   token: session token.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
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

// swagger:operation POST /tables
//
// Serves the client tables for analysis.
// ---
// consumes:
// - text/plain
// produces:
// - text/plain
// parameters:
// - id: user's id
// responses:
//  '200':
//   months: percent diference between the month's salary median and the
//   overall median.
//   hist: number of public servers divided by salary ranges.
//   orgs: best paying public organizations.
//   pos: best paying positions.
//  '401':
//   Result: Unauthorized Access.
//   Code: 9
func (a *App) serveDSTables(w http.ResponseWriter, r *http.Request) {

	cp, _ := os.Getwd()
	path := cp + "/files/"

	files := []string{
		"mostCommon.json",
		"topmonths.json",
		"toporgs.json",
		"toppos.json",
	}

	bestMonths := make(map[string]float64)
	bestOrgs := make(map[string]float64)
	bestPos := make(map[string]float64)
	mostCommon := make(map[string]int)

	for _, file := range files {
		logs.Info("API/serveDSTables", "Opening json file.")

		jsonFile, err := os.Open(path + file)
		if err != nil {
			logs.Errorf("Can't load DS Table.", err.Error())
		}

		logs.Info("API/serveDSTables", "Reading json file.")
		byt, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			logs.Errorf("Can't read DS Table.", err.Error())
		}

		switch file {
		case "topmonths.json":
			logs.Info("API/serveDSTable", "Unmarshaling topmonthsjson.")
			err := json.Unmarshal(byt, &bestMonths)
			if err != nil {
				logs.Errorf("Can't unmarshal DS Table.", err.Error())
			}
		case "toporgs.json":
			logs.Info("API/serveDSTable", "Unmarshaling toporgs.json")
			err := json.Unmarshal(byt, &bestOrgs)
			if err != nil {
				logs.Errorf("Can't unmarshal DS Table.", err.Error())
			}
		case "toppos.json":
			logs.Info("API/serveDSTable", "Unmarshaling toppos.json")
			err := json.Unmarshal(byt, &bestPos)
			if err != nil {
				logs.Errorf("Can't unmarshal DS Table.", err.Error())
			}
		case "mostCommon.json":
			logs.Info("API/serveDSTable", "Unmarshaling mostCommon.json")
			err := json.Unmarshal(byt, &mostCommon)
			if err != nil {
				logs.Errorf("Can't unmarshal DS Table.", err.Error())
			}
		}

		jsonFile.Close()
	}

	type finalDSJson struct {
		Hist   map[string]int     `json:"hist"`
		Months map[string]float64 `json:"months"`
		Orgs   map[string]float64 `json:"orgs"`
		Pos    map[string]float64 `json:"pos"`
	}

	jsonToServe := finalDSJson{
		Hist:   mostCommon,
		Months: bestMonths,
		Orgs:   bestOrgs,
		Pos:    bestPos,
	}

	marsh, err := json.Marshal(jsonToServe)
	if err != nil {
		logs.Errorf("Can't marshal DS Tables.", err.Error())
	}

	w.Write(marsh)
}
