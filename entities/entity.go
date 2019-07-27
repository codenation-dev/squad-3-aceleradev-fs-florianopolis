package entity

import (
	"time"
)

//Cliente entity
type Cliente struct {
	ID            int    `json:"id"`
	Nome          string `json:"nome"`
	IDFuncPublico int    `json:"idfuncpublico"`
}

//FuncPublico entity
type FuncPublico struct {
	ID               int     `json:"id"`
	Nome             string  `json:"nome"`
	Cargo            string  `json:"cargo"`
	Orgao            string  `json:"orgao"`
	Remuneracaodomes float64 `json:"remuneracaodomes"`
	RedutorSalarial  float64 `json:"redutorsalarial"`
	TotalLiquido     float64 `json:"totalliquido"`
	Updated          bool    `json:"updated"`
	ClientedoBanco   bool    `json:"clientedobanco"`
}

//Notificacao entity
type Notificacao struct {
	ID    int              `json:"id"`
	Data  time.Time        `json:"data"`
	Lista NotificacaoLista `json:"lista"`
}

type NotificacaoLista struct {
	ClientesDoBanco         []string `json:"ClientesDoBanco"`
	TopFuncionariosPublicos []string `json:"TopFuncionariosPublicos"`
}

//EmailEnviado entity
type EmailEnviado struct {
	ID            int       `json:"id"`
	IDNotificacao int       `json:"idnotificacao"`
	EmailUsuario  string    `json:"emailusuario"`
	Data          time.Time `json:"data"`
}

//Usuario entity
type Usuario struct {
	ID      int    `json:"id"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Email   string `json:"email"`
}

//Historico entity
type Historico struct {
	ID   int         `json:"id"`
	Data time.Time   `json:"data"`
	JSON FuncPublico `json:"pessoa"`
}

// Credentials struct
type Credentials struct {
	Password string `json:"password"`
	Usermail string `json:"usermail"`
}

type TokenSt struct {
	Token string `json:"token,omitempty"`
}

//MailType References to a Mail
type MailType struct {
	ID   int    `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
	Mail string `json:"Mail,omitempty"`
}

//Resum references to a resum of data
type Resum struct {
	Name string    `json:"Name,omitempty"`
	Date time.Time `json:"Date,omitempty"`
	Data []Data    `json:"DATA,omitempty"`
}

//Data references to a Data
type Data struct {
	Name     string   `json:"Name,omitempty"`
	LineName string   `json:"LineName,omitempty"`
	ColName  string   `json:"ColName,omitempty"`
	Lines    []string `json:"Lines,omitempty"`
	Cols     []string `json:"Cols,omitempty"`
}

type Warn struct{
	ID    int         `json:"id"`
	Data  time.Time   `json:"data"`
	Lista NotificacaoLista `json:"lista"`
	EmailsEnviados []EmailEnviado `json:"emails"`
}

//Mailrequest Type for sending a mail
type Mailrequest struct {
	Subject string `json:"Subject"`
	Targets []Target `json:"Targets"`
	Message string `json:"Message"`
	TopNames []string `json:"TopNames,omitempty"`
	Names []string `json:"Names"`
	Link string `json:"Link"`
	ID int `json:"ID"`
}

//PasswordRequest Type for sending password
type PasswordRequest struct {
	Subject string `json:"Subject"`
	Target Target `json:"Target"`
	Message string `json:"Message"`
}

//Target struct refer email to send
type Target struct {
	Name string `json:"Name"`
	Mail string `json:"Mail"`
	ID string `json:"ID,omitempty"`
}

