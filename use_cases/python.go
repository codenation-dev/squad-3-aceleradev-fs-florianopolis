package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	CreateBestMonthsTable()
}

func CreateBestMonthsTable() {
	cmd := exec.Command("../use_cases/topMonth.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}

func CreateBestOrgsTable() {
	cmd := exec.Command("../use_cases/topOrg.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())

}

func CreateBestPosTable() {
	cmd := exec.Command("../use_cases/topPos.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())

}
