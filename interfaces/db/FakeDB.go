package database

import ("errors"
"golang.org/x/crypto/bcrypt"
"fmt")

//FakeDB To Test
type FakeDB struct {
	User string
	Pass string	
}

//GetPasswordHash Implement Interface 
func (F FakeDB) GetPasswordHash(Usermail string) ([]byte,error) {
	if (Usermail==F.User) {
		return bcrypt.GenerateFromPassword([]byte(F.Pass),2)
	} 
		return nil, errors.New(fmt.Sprintf("No Match, %s and %s",F.User,Usermail))
}