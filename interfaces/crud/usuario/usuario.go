package usuario

import (
	"errors"
	"fmt"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
)

//Insert new Usuario
func Insert(user *entity.Usuario) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Insert(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	if !checkUsuario(user.Email) {
		logs.Info("Insert(Usuario)", "Trying to register new user...")
		//NULLIF Prevents from creating empty string field in table Usuario
		squery := `INSERT INTO USUARIO (usuario, senha, email) VALUES(NULLIF("` + user.Usuario + `",""), NULLIF("` +
			user.Senha + `",""), NULLIF("` + user.Email + `",""));`
		_, erro = dbi.Database.Query(squery)
		if erro != nil {
			logs.Errorf("Insert(Usuario)", erro.Error())
		}
		return erro
	} else {
		logs.Info("Insert(Usuario)", "Email from User already exists")
		return errors.New("Email from User already exists")
	}
}

//Delete Usuario by ID
func Delete(id int) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Delete(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	_, erro = dbi.Database.Query(`DELETE FROM USUARIO WHERE id = ` + strconv.Itoa(id))
	if erro != nil {
		logs.Errorf("Delete(Usuario)", erro.Error())
	}
	return erro
}

//Update Usuario by ID
func Update(user *entity.Usuario) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	logs.Info("Update(Usuario)", "Trying to update user...")
	//NULLIF Prevents from creating empty string field in table Usuario
	squery := `UPDATE USUARIO SET usuario = NULLIF("` + user.Usuario + `",""), senha = NULLIF("` + user.Senha +
		`",""), email = NULLIF("` + user.Email + `","") WHERE id = ` + strconv.Itoa(user.ID)
	fmt.Println(squery)
	_, erro = dbi.Database.Query(squery)
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	} else {
		logs.Info("Update(Usuario)", "User updated successfully!")
	}
	return erro
}

//Get Usuario by ID
func GetUsuarioByID(id int) (*entity.Usuario, error) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("GetUsuarioByID(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE id = ` + strconv.Itoa(id)
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("GetUsuarioByID(Usuario)", erro.Error())
			}
		}
	}
	return &user, erro
}

//Check if user already exists before inserting new user (Get Usuario by Email)
func checkUsuario(email string) bool {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("checkUsuario(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE email = "` + email + `";`
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("checkUsuario(Usuario)", erro.Error())
			}
		}
	}
	if email == user.Email {
		return true
	} else {
		return false
	}
}

//Search user by Mail
func SearchUsuarioByMail(email string) (bool, *entity.Usuario) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("SearchUsuarioByMail(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE email = "` + email + `";`
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("SearchUsuarioByMail(Usuario)", erro.Error())
			}
		}
	}
	if email == user.Email {
		return true, &user
	}
	return false, nil
}

func GetAllMails() []entity.Target {
	dbi, erro := db.Init()
	defer dbi.Database.Close()
	if erro != nil {
		logs.Errorf("GetAllMails(Usuario)", erro.Error())
	}
	seleciona, erro := dbi.Database.Query(`SELECT usuario, email, id FROM USUARIO`)
	if erro != nil {
		logs.Errorf("GetAllMails(Usuario)", erro.Error())
	}
	var mailList []entity.Target
	var onemail entity.Target
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&onemail.Name, &onemail.Mail,&onemail.ID)
			if erro != nil {
				logs.Errorf("GetAllMails(Usuario)", erro.Error())
			}
			mailList = append(mailList, onemail)
		}
	}
	return mailList
}
