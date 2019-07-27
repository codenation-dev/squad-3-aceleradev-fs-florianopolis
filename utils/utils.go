package utils

import (
	"math/rand"
	"time"
  )


 
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  
func GeneratePassword() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))  
	b := make([]byte, 8)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
		return string(b)
}
  
func ConvertDateTimeSQL(date string) time.Time {
	var d time.Time 
	d, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return d
	}
	return d
}