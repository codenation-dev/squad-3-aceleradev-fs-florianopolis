package entity

import (
	"time"
)

//Usuario entity
type Usuario struct {
	ID          int    `json:"id"`
	Cpf         uint64 `json:"cpf"`
	Nome        string `json:"nome"`
	Senha       string `json:"senha"`
	Email       string `json:"email"`
	RecebeEmail bool   `json:"recebeemail"`
}

//Pessoa entity
type Pessoa struct {
	ID                     int     `json:"id"`
	Nome                   string  `json:"nome"`
	Cargo                  string  `json:"cargo"`
	Orgao                  string  `json:"orgao"`
	Remuneracaodomes       float64 `json:"remuneracaodomes"`
	RedutorSalarial        float64 `json:"redutorsalarial"`
	TotalLiquido           float64 `json:"totalliquido"`
	Update                 bool    `json:"update"`
	ClientedoBanco         bool    `json:"clientedobanco"`
}

//Notificacao entity
type Notificacao struct {
	ID       int `json:"id"`
	IDPessoa int `json:"idpessoa"`
}

//NotificacaoUsuario entity
//PRECISAMOS REVISAR ESSA PARTE
type NotificacaoUsuario struct {
	Usuario     Usuario     `json:"idusuario"`
	Notificacao Notificacao `json:"idnotificacao"`
	Data        time.Time   `json:"data"`
}
