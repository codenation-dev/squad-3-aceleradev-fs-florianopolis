# UATI BANK CLIENT/PROSPECT FINDER

## PURPOSE
Finds public servants from the state of São Paulo that classify as a good
potential client, matches them to existing clients, if relevant, and notifies
the user by e-mail.

## OPEN API
To be defined.

## DISCLAIMER
Project being made as a part of an education course on Fullstack Technology by
Codenation.dev. The UATI bank referenced in this project is fictitious.

## DEPENDENCIES
This project uses:
## 1. FrontEnd
  "axios": "^0.19.0", <br />
  "react": "^16.8.6", <br />
  "node.js": "^12.7.0", <br />
  "npm" : "^6.10.2", <br />
  "react-dom": "^16.8.6", <br />
  "react-redux": "^7.1.0", <br />
  "react-redux-es": "^4.4.5",  <br />
  "react-router": "^5.0.1",  <br />
  "react-router-dom": "^5.0.1",  <br />
  "react-scripts": "3.0.1",  <br />
  "recharts": "^1.6.2",  <br />
  "redux": "^4.0.4",  <br />
  "redux-persist": "^5.10.0",  <br />
  "redux-saga": "^1.0.5",  <br />
  "saga": "^4.0.0-alpha"  <br />
## 2.1. BackEnd GO
  "github.com/dgrijalva/jwt-go"  <br />
  "github.com/gorilla/context"  <br />
  "github.com/gorilla/mux"  <br />
  "github.com/gorilla/handlers"  <br />
  "golang.org/x/crypto/bcrypt"  <br />
  "github.com/howeyc/gopass"  <br />
  "github.com/stretchr/testify/assert"  <br />
  "github.com/go-sql-driver/mysql"  <br />
  "github.com/robfig/cron"  <br />
## 2.2. BackEnd Python 3
  mysql.connector  <br />
  numpy  <br />
  collections  <br />
  operator  <br />
## 2.3. MySQL Database
  MySQL 8.0.10 or MariaDB 10.3 or Latest versions
  
## How to install
1) You need to add environment variables into your ~/.bash file or ~/.zshrc  <br />
  export MYSQL_HOST=127.0.0.1  <br />
  export MYSQL_PORT=3306  <br />
  export MYSQL_USER={YourUserDbName}  <br />
  export MYSQL_PASSWORD={YourUserDbPassword}  <br />
  export MYSQL_DATABASE={YourDbName}  <br />
  export LOG_PATH={YourLogPath}  <br />
1.1) To use email service, you need to create a folder called ".cfg", inside that folder you must to create a file called: credentials
The full path will be look like this: squad-3-aceleradev-fs-florianopolis/services/MailSender/.cfg/credentials
1.2) Inside the credentials, add your email account and your password separated by comma, like this: test@gmail.com,1234

2) After add enviroment variables, execute this query to create your tables:
<pre>
/*SUGGESTION NAME FOR YOUR DATABASE*/
CREATE DATABASE bancouati; 
USE bancouati;

CREATE TABLE `FUNCPUBLICO` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(100) NOT NULL,
  `cargo` varchar(100) NOT NULL,
  `orgao` varchar(45) NOT NULL,
  `remuneracaodomes` double NOT NULL,
  `redutorsalarial` double DEFAULT NULL,
  `totalliquido` double DEFAULT NULL,
  `updated` tinyint(1) DEFAULT '0',
  `clientedobanco` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
);

CREATE TABLE `USUARIO` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `usuario` varchar(30) NOT NULL,
  `senha` char(60) NOT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE INDEX idx_email
ON `USUARIO` (`email`);

CREATE TABLE `NOTIFICACAO` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data` datetime NOT NULL,
  `lista` json NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `HISTORICO` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data` datetime NOT NULL,
  `json` json NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `EMAILENVIADO` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `idnotificacao` INT NOT NULL,
  `emailusuario` VARCHAR(100) NOT NULL,
  `data` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`idnotificacao`) REFERENCES NOTIFICACAO(`id`)
);
</pre>

3) Execute git clone or download this project <br />
Git clone https://github.com/codenation-dev/squad-3-aceleradev-fs-florianopolis.git <br />
4) Enter in the project folder on your local machine <br />
5) Run ./install.sh to install all dependecies from the application <br />
6) Run ./run.sh to run the appliation, it will open in your browser <br />
7) If you want to stop all application run ./stopall.sh to stop the application <br />
8) If you want change the source code, you need to delete the executables files, for this, run ./stopall.sh -d <br />
To build the executables files again, execute ./run.sh again

