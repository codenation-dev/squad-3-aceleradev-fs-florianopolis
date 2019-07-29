package emailenviado

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

}

func TestGetAll(t *testing.T) {
	id := 1 //at least one data in notifications table (id auto increment)
	emailsList := GetAll(id)
	assert.Greater(t, len(emailsList), 0)
}
