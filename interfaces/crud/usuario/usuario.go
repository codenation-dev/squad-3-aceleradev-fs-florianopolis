package usuario

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"strconv"
)

//Insert new Usuario
func Insert(user *entity.Usuario) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Insert(User)", erro.Error())
	}
	defer dbi.Database.Close()
	if !checkUsuario(user.Email) {
		logs.Info("Insert(User)", "Inserting new user in DB...")
		squery := "INSERT INTO USUARIO (usuario, senha, email) VALUES('" +
			user.Usuario + "', '" + user.Senha + "','" + user.Email + "');"
		erro = dbi.ExecQuery(squery)
		return erro
	} else {
		logs.Info("Insert(User)", "User already exists")
		return erro
	}
}

//Delete Usuario by ID
func Delete(id int) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Delete(User)", erro.Error())
	}
	defer dbi.Database.Close()
	erro = dbi.ExecQuery("DELETE FROM USUARIO WHERE id = " + strconv.Itoa(id))
	return erro
}

//Update Usuario by ID
func Update(user *entity.Usuario) error {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("Update(User)", erro.Error())
	}
	defer dbi.Database.Close()
	squery := "UPDATE USUARIO SET usuario = '" + user.Usuario + "', senha = '" + user.Senha +
		"', email = '" + user.Email + "' WHERE id = " + strconv.Itoa(user.ID)
	erro = dbi.ExecQuery(squery)
	return erro
}

//Get Usuario by ID
func GetUsuarioByID(id int) (*entity.Usuario, error) {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("getUsuarioByID", erro.Error())
	}
	defer dbi.Database.Close()
	squery := "SELECT * FROM USUARIO WHERE id = " + strconv.Itoa(id)
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario

	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("GetUsuarioByID", erro.Error())
			}
		}
	}
	return &user, erro
}

//Check if user already exists before inserting new user (Get Usuario by Email)
func checkUsuario(email string) bool {
	dbi, erro := db.Init()
	if erro != nil {
		logs.Errorf("checkUsuario", erro.Error())
	}
	defer dbi.Database.Close()
	squery := "SELECT * FROM USUARIO WHERE email = '" + email + "';"
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("checkUsuario", erro.Error())
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
func SearchUsuarioByMail(email string) (bool ,*entity.Usuario) {
	
	dbi, erro := db.Init()
	
	if erro != nil {
		logs.Errorf("checkUsuario", erro.Error())
	}
	
	defer dbi.Database.Close()
	squery := "SELECT * FROM USUARIO WHERE email = '" + email + "';"
	seleciona, erro := dbi.Database.Query(squery)
	var user entity.Usuario
	if erro == nil {
		for seleciona.Next() {
			erro := seleciona.Scan(&user.ID, &user.Usuario, &user.Senha, &user.Email)
			if erro != nil {
				logs.Errorf("checkUsuario", erro.Error())
			}
		}
	}
	
	if email == user.Email {
		return true, &user
	}
	
	return false, nil
}