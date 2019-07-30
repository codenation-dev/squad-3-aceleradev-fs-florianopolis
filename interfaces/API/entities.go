package api

import ("squad-3-aceleradev-fs-florianopolis/entities"
"crypto/rsa"
"github.com/gorilla/mux"
"time")

const (
	//Success Code for response 
	Success = 1
	//Empty Code for response
	Empty 	= 2
	//Error Code for response
	Error   = 3
	//Unauth Code for response
	Unauth = 9
	
	changesSave = 12

	fileUploadedSuccess = 13

	userRemoveSuccess = 14

 )
 

//App the struct for the app
type App struct {
	Router    *mux.Router
	Database  string
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

//DataEmailUsuario define json struct
type DataEmailUsuario struct{
	ID	  int		 `json:"id"`
	Data  time.Time   `json:"data"`
}

type passT struct {
	Subject string `json:"Subject"`
	Message string `json:"Message"`
	Target entity.Target `json:"Target"`
}

//Result pattern response for all requests
type Result struct {
	Result    string   `json:"Result"`
	Code	  int	   `json:"Code,omitempty"`	
	Token     string   `json:"token,omitempty"`
	Warn      *Warn     `json:"Warn,omitempty"`
	Mail      *entity.Target `json:"Mail,omitempty"`
	DataResum *Resum    `json:"DataResum,omitempty"`
	Usermails *[]entity.Target `json:"UsermailList,omitempty"`
}

//ListaClientes define json struct
type ListaClientes struct {
	Nome string `json:"nome,omitempty"`
}

type Resum entity.Resum
type tokenSt entity.TokenSt
type credentials entity.Credentials
type Warn entity.Warn