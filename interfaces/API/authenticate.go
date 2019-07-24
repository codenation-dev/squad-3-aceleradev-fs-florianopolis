package api

import (jwt "github.com/dgrijalva/jwt-go"
"io/ioutil"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"time"
"golang.org/x/crypto/bcrypt"
"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario")

const (
	pubPath="keys/app.rsa.pub"
	privPath="keys/app.rsa"
)

//initKey initialize key from file
func (a *App) initKey() {
	
	signKey, err := ioutil.ReadFile(privPath)
	a.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKey)

	if (err!=nil){
		logs.Errorf("App_initKey_SignKey",err.Error())
		panic(err)
	}

	verifyKey, err := ioutil.ReadFile(pubPath)
	a.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKey)
	
	if (err!=nil){
		logs.Errorf("App_initKey_VerifyKey",err.Error())
		panic(err)
	}
	logs.Info("App_initKey","Sucesso na inilizaçao do par de chaves")
}

//GenerateJWT Generate Token
func (a *App) GenerateJWT(Username string) (string,error) {
	
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = Username
	claims["exp"] = time.Now().Add(time.Minute *2).Unix()

	tokenString,err := token.SignedString(a.signKey)

	if(err!=nil){
		logs.Errorf("App_GenerateJWT",err.Error())
		return "",nil
	}

	return tokenString, nil
}

func (a *App) tryLogin(c *credentials) bool {

	valid, usr := usuario.SearchUsuarioByMail()
	
	Hash, err := GetPassword(c.Usermail)
	
	if(err != nil){
		logs.Errorf("App_loginAttempt_User", err.Error())
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(Hash),[]byte(c.Password))

	if(err != nil){
		logs.Errorf("App_loginAttempt_Pass", err.Error())
		return false
	}
	
	return true
	}
}

func (a *App) tokenVerify(T tokenSt) (*jwt.Token,error) {
	return jwt.Parse(T.Token,func(token *jwt.Token) (interface{},error){return a.verifyKey,nil})
}