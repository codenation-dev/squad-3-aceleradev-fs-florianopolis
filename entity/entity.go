package entity

type Usuario struct{
	ID int `json:id`
	Cpf int64 `json:cpf`
	Nome string `json:nome`
	Senha string `json:senha`
	Email string `json:email`
	FuncionarioPuplico bool `json:funcionariopublico`
}

/*
CREATE DATABASE `bancouati`;
CREATE TABLE `bancouati`.`usuario` (
	`id` `id` INT(11) NOT NULL AUTO_INCREMENT,
	`cpf` BIGINT NOT NULL,
	`nome` VARCHAR(100) NOT NULL,
	`senha` VARCHAR(100) NOT NULL,
	`email` VARCHAR(100) NOT NULL,
	`funcionariopublico` TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`, `cpf`));

INSERT INTO `bancouati`.`usuario` (`cpf`, `nome`, `senha`, `email`, `funcionariopublico`) VALUES ('33333333333', 'Joao da Silva', '12345', 't@t.com', '0');
  */
