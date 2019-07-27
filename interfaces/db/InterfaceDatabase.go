package database

// InterfaceDatabase interface for implementing DB
type InterfaceDatabase interface {
	GetPasswordHash(Usermail string) ([]byte,error) 
}