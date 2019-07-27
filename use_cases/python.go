package python

import (
	"os"
	"os/exec"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
)

// Runs python script that generates the topmonths.json file containing a table
// of the months in which the public servers earned more.
func CreateBestMonthsTable() {
	cmd := exec.Command("../use_cases/topMonth.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logs.Info(cmd.Run())
}

// Runs python script that generates the toporgs.json file containing a table
// of the organizations that pay the most.
func CreateBestOrgsTable() {
	cmd := exec.Command("../use_cases/topOrg.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logs.Info(cmd.Run())

}

// Runs python script that generates the toppos.json file containing a table of
// the positions that earn the most.
func CreateBestPosTable() {
	cmd := exec.Command("../use_cases/topPos.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logs.Info(cmd.Run())

}
