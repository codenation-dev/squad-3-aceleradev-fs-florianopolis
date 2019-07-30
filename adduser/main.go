package main

import (
    "fmt"
	"golang.org/x/crypto/bcrypt"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/interfaces/crud/usuario"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"regexp"
	"github.com/howeyc/gopass"
)

func main()  {
	if !UserExists() {
		dbi, err := db.Init()
		defer dbi.Database.Close()	
		var user entity.Usuario
		for {
			fmt.Print("Inform your full name: ")
			_, err:= fmt.Scanln(&user.Usuario)
			if err == nil {
				break
			}else{
				fmt.Print(err.Error())
			}
		}
		for {
			fmt.Print("Inform your email: ")
			_, err := fmt.Scanln(&user.Email)
			if validateMailType(&user){
				if err == nil {
					break
				}else{
					fmt.Print(err.Error())
				}
			}else{
				fmt.Println("email invalid. try again")
			}
		}
		for {
			fmt.Print("Inform your password: ")
			maskedPassword1, _ := gopass.GetPasswdMasked()
			fmt.Print("Inform your password again: ")
			maskedPassword2, _ := gopass.GetPasswdMasked()
			if string(maskedPassword1) == string(maskedPassword2) {
				pass, err := bcrypt.GenerateFromPassword([]byte(user.Senha), 10)
				user.Senha = string(pass)
				if err == nil {
					break
				}else{
					fmt.Print(err.Error())
				}
			}else{
				fmt.Println("Password does not match. Try again!")
			}

		}
		if err != nil {
			fmt.Print(err.Error())
		}
		err = usuario.Insert(&user, dbi)
		if err != nil {
			fmt.Print(err.Error())
		} 
	}
	
}

func validateMailType(u *entity.Usuario) bool {
	validateEmail := regexp.MustCompile(`^[-a-z0-9~!$%^&*_=+}{\'?]+(\.[-a-z0-9~!$%^&*_=+}{\'?]+)*@([a-z0-9_][-a-z0-9_]*(\.[-a-z0-9_]+)*\.(aero|arpa|biz|com|coop|edu|gov|info|int|mil|museum|name|net|org|pro|travel|mobi|[a-z][a-z])|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,5})?$`)
	return validateEmail.MatchString(u.Email)
}

func UserExists() bool {
	dbi, err := db.Init()
	defer dbi.Database.Close()
	seleciona, erro := dbi.Database.Query(`SELECT id FROM USUARIO LIMIT 1`)
	if erro != nil {
		fmt.Print(err.Error())
	}
	var id int
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&id)
			if erro != nil {
				fmt.Print(err.Error())
			}
		}
	}
	return id > 0
}