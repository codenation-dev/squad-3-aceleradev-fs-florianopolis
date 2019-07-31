package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfoFromCSVFile(t *testing.T) {

}

func TestProcessMultiLinesCSVFile(t *testing.T) {

}

func TestPersistLinesInDb(t *testing.T) {

}

func TestOpenCSV(t *testing.T) {

}

func TestAfterProcess(t *testing.T) {

}

func TestOpenCSVAndInsertCSV(t *testing.T) {

}

func TestCheckPersonInDB(t *testing.T) {

}

func TestIsClient(t *testing.T) {
	name := "fakename"
	isNameClient := isClient(name)
	assert.Equal(t, false, isNameClient)
}

func TestPersistPessoa(t *testing.T) {

}

func TestChangeComma(t *testing.T) {

}

func TestGetHashFromFile(t *testing.T) {

}

func TestGetLastFiles(t *testing.T) {

}
