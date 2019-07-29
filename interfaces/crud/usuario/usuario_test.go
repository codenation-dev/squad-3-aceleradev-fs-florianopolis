package usuario

import (
	entity "squad-3-aceleradev-fs-florianopolis/entities"
	db "squad-3-aceleradev-fs-florianopolis/interfaces/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	// Creates sqlmock database connection and a mock to manage expectations
	dbm, mock, err := sqlmock.New()
	//dbm, mock, err := New()
	defer dbm.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO USUARIO").WithArgs("test", "test", "test@test.com").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	//mock.ExpectExec("INSERT INTO USUARIO").WithArgs("test", "test", "test@test.com").WillReturnResult(sqlmock.NewResult(1, 1))

	User := entity.Usuario{
		Usuario: "test",
		Email:   "test@test.com",
		Senha:   "test"}

	var dbs db.MySQLDatabase
	dbs.Database = dbm

	// now we execute our method
	if err = Insert(&User, &dbs); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	/*if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}*/
}

func TestUpdate(t *testing.T) {
	//needs mocking! crud
}

func TestDelete(t *testing.T) {
	//needs mocking! crud
}

func TestGetUsuarioByID(t *testing.T) {
	var user *entity.Usuario
	//application has at least one user
	user, err := GetUsuarioByID(1)
	assert.Equal(t, user.ID, 1)
	assert.Nil(t, err)
}

func TestCheckUsuario(t *testing.T) {
	email := "invalidemailgmail.com" //invalid email without @
	existingUser, erro := CheckUsuario(email)
	assert.Equal(t, false, existingUser)
	assert.Nil(t, erro)
}

func TestSearchUsuarioByMail(t *testing.T) {
	email := "invalidemailgmail.com" //invalid email without @
	existingUser, err := SearchUsuarioByMail(email)
	assert.Nil(t, err)
	assert.Equal(t, false, existingUser)
}

func TestGetAllMails(t *testing.T) {
	var allMailsList []entity.Target
	allMailsList = GetAllMails()
	assert.Greater(t, len(allMailsList), 0)
}
