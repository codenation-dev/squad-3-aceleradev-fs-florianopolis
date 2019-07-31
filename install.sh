#!/bin/bash
path=`pwd`
reactpath="/frontend/"

echo "YOU NEED TO HAVE GO, PYTHON 3, PIP3, NODEJS AND NPM INSTALED ON YOU MACHINE"
echo "IF YOU DON'T HAVE ONE OF THIS APP INSTALLED, PLEASE CANCEL THIS INSTALATION"
echo "Do you want continue? write (y) to yes or (n) to no: "
read yesno < /dev/tty
if [ "x$yesno" = "xy" ];then
    echo "Instaling Go dependecies..."
    go get github.com/dgrijalva/jwt-go
    go get github.com/gorilla/context
    go get github.com/gorilla/mux
    go get github.com/gorilla/handlers
    go get golang.org/x/crypto/bcrypt
    go get github.com/howeyc/gopass
    go get github.com/stretchr/testify/assert
    go get github.com/go-sql-driver/mysql
    go getgithub.com/robfig/cron
    echo "Instaling Python 3 dependecies..."
    pip3 install mysql-connector
    pip3 install numpy
    pip3 install collections
    pip3 install operator
    echo "Instaling React dependecies..."
    npm install
else
   echo "Cancel, nothing was done"
fi