package usuario

import (
	"errors"
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
)

//Insert New Usuario in DB
func Insert(user *entity.Usuario, dbi *db.MySQLDatabase) error {
	tx, err := dbi.Database.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	/*dbi, erro := db.Init() //changed to mock DB for unit test
	if erro != nil {
		logs.Errorf("Insert(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()*/
	userAlreadyExists, erro := CheckUsuario(user.Email)
	if erro != nil {
		logs.Errorf("Insert(Usuario)", erro.Error())
	}
	if !userAlreadyExists {
		logs.Info("Insert(Usuario)", "Trying to register new ur...")
		//NULLIF Prevents from creating empty string field in table Usuariose
		/*squery := `INSERT INTO USUARIO (usuario, senha, email) VALUES(NULLIF("` + user.Usuario + `",""), NULLIF("` +
		user.Senha + `",""), NULLIF("` + user.Email + `",""));`*/
		//result, erro := dbi.Database.Query(squery)
		_, erro := tx.Exec(`INSERT INTO USUARIO (usuario, senha, email) VALUES(NULLIF(?,""), NULLIF(?,""), NULLIF(?,""));`, user.Usuario, user.Senha, user.Email)
		//defer result.Close()
		if erro != nil {
			logs.Errorf("Insert(Usuario)", erro.Error())
		}
		return erro
	}
	logs.Info("Insert(Usuario)", "Email from User already exists")
	return errors.New("Email from User already exists")
}

//Delete Usuario by ID
func Delete(id int, dbi *db.MySQLDatabase) error {
	tx, err := dbi.Database.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	/*dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Delete(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	result, erro := dbi.Database.Query(`DELETE FROM USUARIO WHERE id = ` + strconv.Itoa(id))
	defer result.Close()*/
	_, erro := tx.Exec(`DELETE FROM USUARIO WHERE id = ?`, id)
	if erro != nil {
		logs.Errorf("Delete(Usuario)", erro.Error())
	}
	return erro
}

//Update Usuario by ID
func Update(user *entity.Usuario, dbi *db.MySQLDatabase) error {
	tx, err := dbi.Database.Begin()
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	/*dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()*/
	logs.Info("Update(Usuario)", "Trying to update user...")
	//NULLIF Prevents from creating empty string field in table Usuario
	/*squery := `UPDATE USUARIO SET usuario = NULLIF("` + user.Usuario + `",""), senha = NULLIF("` + user.Senha +
		`",""), email = NULLIF("` + user.Email + `","") WHERE id = ` + strconv.Itoa(user.ID)
	result, erro := dbi.Database.Query(squery)
	defer result.Close()*/
	_, erro := tx.Exec(`UPDATE USUARIO SET usuario = NULLIF(?,""), senha = NULLIF(?,""), email = NULLIF(?,"") WHERE id = ?`,
		user.Usuario, user.Senha, user.Email, user.ID)
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	} else {
		logs.Info("Update(Usuario)", "User updated successfully!")
	}
	return erro
}

//UpdateWithoutPass Usuario by ID
func UpdateWithoutPass(user *entity.Usuario) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	logs.Info("Update(Usuario)", "Trying to update user...")
	//NULLIF Prevents from creating empty string field in table Usuario
	squery := `UPDATE USUARIO SET usuario = NULLIF("` + user.Usuario + `",""), email = NULLIF("` + user.Email + `","") WHERE id = ` + strconv.Itoa(user.ID)
	result, erro := dbi.Database.Query(squery)
	defer result.Close()
	if erro != nil {
		logs.Errorf("Update(Usuario)", erro.Error())
	} else {
		logs.Info("Update(Usuario)", "User updated successfully!")
	}
	return erro
}

//GetUsuarioByID Usuario by ID
func GetUsuarioByID(id int) (*entity.Usuario, error) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("GetUsuarioByID(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE id = ` + strconv.Itoa(id)
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
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

//CheckUsuario checks if user already exists before inserting new user (Get Usuario by Email)
func CheckUsuario(email string) (bool, error) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("CheckUsuario(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE email = "` + email + `";`
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("CheckUsuario(Usuario)", erro.Error())
			}
		}
	}
	if email == user.Email {
		return true, erro
	}
	return false, erro
}

//SearchUsuarioByMail user by Mail
func SearchUsuarioByMail(email string) (bool, *entity.Usuario) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("SearchUsuarioByMail(Usuario)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := `SELECT * FROM USUARIO WHERE email = "` + email + `";`
	seleciona, erro := dbi.Database.Query(squery)
	defer seleciona.Close()
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

//GetAllMails get all user mails
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
			erro := seleciona.Scan(&onemail.Name, &onemail.Mail, &onemail.ID)
			if erro != nil {
				logs.Errorf("GetAllMails(Usuario)", erro.Error())
			}
			mailList = append(mailList, onemail)
		}
	}
	return mailList
}
