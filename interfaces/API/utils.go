package api

import (
	"math/rand"
	"time"
  )

func validateMailType(m MailType) bool {
	return true
}
 
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  
func generatePassword() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))  
	b := make([]byte, 8)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
		return string(b)
}
  