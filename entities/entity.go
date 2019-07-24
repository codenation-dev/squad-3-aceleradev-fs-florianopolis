package entity

import "time"

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
	ID    int         `json:"id"`
	Data  time.Time   `json:"data"`
	Lista FuncPublico `json:"pessoa"`
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
	ID   int  		 `json:"id"`
	Data time.Time 	 `json:"data"`
	JSON FuncPublico `json:"pessoa"`
}
