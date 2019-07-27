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
