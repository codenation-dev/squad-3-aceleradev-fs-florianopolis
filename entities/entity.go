package entity

import (
	"time"
)
//Usuario entity
type Usuario struct{
	ID int `json:id`
  	Cpf uint64 `json:cpf`
  	Nome string `json:nome`
  	Senha string `json:senha`
  	Email string `json:email`
  	Recebeemail bool `json:recebeemail`
}

//Pessoa entity
type Pessoa struct{
	ID int `json:id`
	IdArquivoTransparencia int `json:idarquivotransparencia`
	Nome string `json:nome`
	Cargo string `json:cargo`
	Orgao string `json:orgao`
	Remuneracaodomes float64 `json:remuneracaodomes`
	RedutorSalaria float64 `json:redutorsalarial`
	TotalLiquido float64 `json:totalliquido`
	Update bool `json:update`
	ClientedoBanco bool`json:clientedobanco`
}

//Notificacao entity
type Notificacao struct{
	Id int `json:id`
	IdPessoa int `json:idpessoa`
}

//NotificacaoUsuario entity
type NotificacaoUsuario struct{
	IdUsuario int `json:idusuario`
	IdNotificacao int `json:idnotificacao`
	Data time.Time `json:data`
}