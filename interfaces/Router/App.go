package Router

import ("github.com/codenation-dev/squad-3-aceleradev-fs-florianopolis/entities/logs"
		"github.com/codenation-dev/squad-3-aceleradev-fs-florianopolis/interfaces/db"
		"github.com/gorilla/mux"
		jwt "github.com/dgrijalva/jwt-go"
		"net/http"
		"golang.org/x/crypto/bcrypt"
		"fmt"
		"time"
		"os"
	)

// App references Dependecies
type App struct {
	Router *mux.Router
	Database database.InterfaceDatabase
}

// Credentials struct
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func GenerateJWT(Username string) (string,error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorize"] = true
	claims["user"] = Username
	claims["exp"] = time.Now().Add(time.Minute *5).Unix()

	tokenString,err := token.SignedString(os.Getenv("JWT_TOKEN"))

	if(err!=nil){
		logs.Errorf("GenerateJWT Error",fmt.Sprintf("%s",err))
		return "",nil
	}

	return tokenString, nil
}

func (a *App) loginAttempt(c *Credentials) bool {
	
	Hash, err := a.Database.GetPasswordHash(c.Username)
	
	if(err != nil){
		logs.Errorf("App Login Attempt", fmt.Sprintf("%s",err))
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(Hash),[]byte(c.Password))

	if(err != nil){
		logs.Errorf("App Login Attempt", fmt.Sprintf("%s",err))
		return false
	}
	
	return true

}

// LoginHandler handle the login
func (a *App) LoginHandler(w http.ResponseWriter,r *http.Request) {

	vars := mux.Vars(r)
	cred := &Credentials{
		Password:vars["pwd"],
		Username:vars["user"]}
	
		if(a.loginAttempt(cred)){

	}

}

func (a *App) MainMiddleware(w http.ResponseWriter, r *http.Request) {

}

