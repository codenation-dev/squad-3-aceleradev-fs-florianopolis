package api

import (
	"crypto/rsa"
	"time"

	"github.com/gorilla/mux"
)

type Result struct {
	Result    string   `json:"Result"`
	Token     string   `json:"Token,omitempty"`
	Warn      Warn     `json:"Warn,omitempty"`
	Warns     WarnList `json:"WarnList,omitempty"`
	Mail      MailType `json:"Mail,omitempty"`
	Mails     MailList `json:"Mails,omitempty"`
	DataResum Resum    `json:"DataResum,omitempty"`
}

type App struct {
	Router    *mux.Router
	Database  string
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

// Credentials struct
type credentials struct {
	Password string `json:"password"`
	Usermail string `json:"usermail"`
}

type tokenSt struct {
	Token string `json:"token"`
}

//Warn References to a Warn
type Warn struct {
	Name   string `json:"Name"`
	Change string `json:"Changes"`
}

//WarnList References to a list of Warns
type WarnList struct {
	Warns []Warn    `json:"Warns"`
	Date  time.Time `json:"Date"`
}

//MailList References to a list of Mails
type MailList struct {
	Mails []MailType `json:"Mails"`
}

//MailType References to a Mail
type MailType struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

//Resum references to a resum of data
type Resum struct {
	Name string    `json:"Name"`
	Date time.Time `json:"Date"`
	Data []Data    `json:"DATA"`
}

//Data references to a Data
type Data struct {
	Name     string   `json:"Name"`
	LineName string   `json:"LineName"`
	ColName  string   `json:"ColName"`
	Lines    []string `json:"Lines"`
	Cols     []string `json:"Cols"`
}

//ListaClientes define json struct
type ListaClientes struct {
	Nome string `json:"nome"`
}
