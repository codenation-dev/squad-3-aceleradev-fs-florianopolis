# UATI CLIENT/PROSPECT FINDER

## PURPOSE
Finds public workers from the state of São Paulo that classify as a good
potential client, matches them to existing clients, if relevant, and notifies
the user by e-mail.

## OPEN API
To define.

## DISCLAIMER
Project being made as a part of an education course on Fullstack Technology by
Codenation.dev. The UATI bank referenced in this project is fictitious.

## DEPENDENCIES
This project use:
## 1. FrontEnd
  "axios": "^0.19.0",
  "react": "^16.8.6",
  "react-dom": "^16.8.6",
  "react-redux": "^7.1.0",
  "react-redux-es": "^4.4.5",
  "react-router": "^5.0.1",
  "react-router-dom": "^5.0.1",
  "react-scripts": "3.0.1",
  "recharts": "^1.6.2",
  "redux": "^4.0.4",
  "redux-persist": "^5.10.0",
  "redux-saga": "^1.0.5",
  "saga": "^4.0.0-alpha"
## 2.1. BackEnd GO
  "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
  "github.com/gorilla/handlers"
	"golang.org/x/crypto/bcrypt"
  "github.com/howeyc/gopass"
  "github.com/stretchr/testify/assert"
  "github.com/go-sql-driver/mysql"
  "github.com/robfig/cron"
## 2.2. BackEnd Python 3
  mysql.connector
  numpy
  collections
  operator
## 2.3. DataBase Mysql
  MySQL 8.0.10 or MariaDB 10.3 or Latest versions
  
## How to install
Add environment variables into your ~/.bash file or ~/.zshrc 
  export MYSQL_HOST=127.0.0.1
  export MYSQL_PORT=3306
  export MYSQL_USER={YourUserDbName}
  export MYSQL_PASSWORD={YourUserDbPassword}
  export MYSQL_DATABASE={YourDbName}
  export LOG_PATH={YourLogPath}
