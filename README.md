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
Add environment variables into your ~/.bash file or ~/.zshrc  <br />
  export MYSQL_HOST=127.0.0.1  <br />
  export MYSQL_PORT=3306  <br />
  export MYSQL_USER={YourUserDbName}  <br />
  export MYSQL_PASSWORD={YourUserDbPassword}  <br />
  export MYSQL_DATABASE={YourDbName}  <br />
  export LOG_PATH={YourLogPath}  <br />
Git clone https://github.com/codenation-dev/squad-3-aceleradev-fs-florianopolis.git <br />
Open the project folder on your local machine <br />
Run ./install.sh to install the application <br />
Run ./run.sh to run the appliation, it will open in your browser <br />
Run ./stopall.sh to stop the application <br />
Run ./stopall.sh -d if you want to delete executable files <br />

