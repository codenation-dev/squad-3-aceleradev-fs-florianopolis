package api

import (
	jwt "github.com/dgrijalva/jwt-go"
	"regexp"
	"io/ioutil"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"time"
	"golang.org/x/crypto/bcrypt"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	"squad-3-aceleradev-fs-florianopolis/entities"
	"fmt"
)

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

//GetToken Generates a token using the received token
func (a *App) GetToken(token *jwt.Token) (string) {
	
	Claims := token.Claims.(jwt.MapClaims)
	R,err := a.GenerateJWT(Claims["user"].(string))
	if err != nil {
		return "tHERE"
	}
	return R
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

	valid, usr := usuario.SearchUsuarioByMail(c.Usermail)
	
	if(!valid){
		logs.Errorf("App_loginAttempt_User", fmt.Sprintf("Cant Get User %s",c.Usermail))
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(usr.Senha),[]byte(c.Password))

	if(err != nil){
		logs.Errorf("App_loginAttempt_Pass", err.Error())
		return false
	}
	
	return true
}

func (a *App) tokenVerify(T tokenSt) (*jwt.Token,error) {
	return jwt.Parse(T.Token,func(token *jwt.Token) (interface{},error){return a.verifyKey,nil})
}

func validateMailType(m entity.Target) bool {
	validateEmail := regexp.MustCompile(`^[-a-z0-9~!$%^&*_=+}{\'?]+(\.[-a-z0-9~!$%^&*_=+}{\'?]+)*@([a-z0-9_][-a-z0-9_]*(\.[-a-z0-9_]+)*\.(aero|arpa|biz|com|coop|edu|gov|info|int|mil|museum|name|net|org|pro|travel|mobi|[a-z][a-z])|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,5})?$`)
	return validateEmail.MatchString(m.Mail)
}