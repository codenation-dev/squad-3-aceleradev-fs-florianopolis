package router

import ("squad-3-aceleradev-fs-florianopolis/entities/logs"
		"squad-3-aceleradev-fs-florianopolis/interfaces/db"
		"github.com/gorilla/mux"
		jwt "github.com/dgrijalva/jwt-go"
		"net/http"
		"golang.org/x/crypto/bcrypt"
		"fmt"
		"time"
		"io/ioutil"
		"encoding/json"
		"crypto/rsa"
)

const (
	pubPath="keys/app.rsa.pub"
	privPath="keys/app.rsa"
)

// App references Dependecies
type App struct {
	Router *mux.Router
	Database database.InterfaceDatabase
	SignKey *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
}

// Credentials struct
type Credentials struct {
	Password string `json:"password"`
	Usermail string `json:"usermail"`
	
}

type TokenSt struct {
	Token string `json:"token"`
}

//StartR Creates a Router
func StartR(Database database.InterfaceDatabase) *App {
	m := mux.NewRouter()
	AppL := &App{Router: m,Database:Database}
	AppL.initKey()
	AppL.setupHandlers()
	return AppL
}

//SetupHandlers Setup the config
func (a *App) setupHandlers() {
	a.Router.HandleFunc("/verify",a.MainMiddleware(a.Verify)).Methods("POST")
	a.Router.HandleFunc("/login",a.LoginHandler).Methods("POST")
	a.Router.HandleFunc("/dashboard",a.MainMiddleware(a.NotImplemented)).Methods("GET")
	a.Router.HandleFunc("/load",a.MainMiddleware(a.NotImplemented)).Methods("POST")
	a.Router.HandleFunc("/register",a.MainMiddleware(a.NotImplemented)).Methods("POST")
}

//initKey initialize key from file
func (a *App) initKey() {
	
	SignKey, err := ioutil.ReadFile(privPath)
	a.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(SignKey)

	if (err!=nil){
		logs.Errorf("App_initKey_SignKey",err.Error())
		panic(err)
	}

	VerifyKey, err := ioutil.ReadFile(pubPath)
	a.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(VerifyKey)
	if (err!=nil){
		logs.Errorf("App_initKey_VerifyKey",err.Error())
		panic(err)
	}
}

//GenerateJWT Generate Token
func (a *App) GenerateJWT(Username string) (string,error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = Username
	claims["exp"] = time.Now().Add(time.Minute *5).Unix()


	tokenString,err := token.SignedString(a.SignKey)

	if(err!=nil){
		logs.Errorf("App_GenerateJWT",err.Error())
		return "",nil
	}

	return tokenString, nil
}

func (a *App) loginAttempt(c *Credentials) bool {
	
	Hash, err := a.Database.GetPasswordHash(c.Usermail)
	
	if(err != nil){
		logs.Errorf("App_loginAttempt_User", fmt.Sprintf("%s",err))
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(Hash),[]byte(c.Password))

	if(err != nil){
		logs.Errorf("App_loginAttempt_Pass", fmt.Sprintf("%s",err))
		return false
	}
	
	return true

}

// LoginHandler handle the login
func (a *App) LoginHandler(w http.ResponseWriter,r *http.Request) {

	decode := json.NewDecoder(r.Body)
	var cred Credentials
	_ = decode.Decode(&cred)
	
	if(a.loginAttempt(&cred)){
			T,E := a.GenerateJWT(cred.Usermail)
			if (E!=nil){
				logs.Errorf("App/Cant create JWT Token",E.Error())
			}
			fmt.Fprintf(w,fmt.Sprintf("this is your token: %s",string(T)))
	} else {
		fmt.Fprintf(w,"Login Error")
	}

}
// r.Method switch(Post/get)
// MainMiddleware Token Verify and Authenticate
func (a *App) MainMiddleware(next http.HandlerFunc) http.HandlerFunc {	
	return func(w http.ResponseWriter, r *http.Request){
	decode := json.NewDecoder(r.Body)
	var TempT TokenSt
	_ = decode.Decode(&TempT)  
	token,err := jwt.Parse(TempT.Token,func(token *jwt.Token) (interface{},error){return a.VerifyKey,nil})
	
	if (err == nil) {
		if (token.Valid) {
			next(w,r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w,"Invalid Token")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w,"Unauthorized Request")
	}
}
}

//Verify Screen Success test
func (a *App) Verify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Success!")
}

//HandlerUpCSV Up CSV
func (a *App) HandlerUpCSV(w http.ResponseWriter, r *http.Request) {
	
}

//HandlerRegisterUser register new user
func (a *App) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {

}

//HandlerDashBoard dash 
func (a *App) HandlerDashBoard(w http.ResponseWriter, r *http.Request) {

}

//HandlerReport report
func (a *App) HandlerReport(w http.ResponseWriter, r *http.Request) {

}

//NotImplemented not implemented
func (a *App) NotImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"NotImplemented But Success")
}