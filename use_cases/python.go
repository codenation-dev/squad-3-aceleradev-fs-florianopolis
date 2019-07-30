package python

import (
	"os"
	"os/exec"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	"sync"
)

// Runs python script that generates the topmonths.json file containing a table
// of the months in which the public servers earned more.
func CreateBestMonthsTable(wg *sync.WaitGroup) {
	logs.Info("Python", "Starting topMonth.py script.")

	cmd := exec.Command("../../use_cases/topMonth.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logs.Errorf("Python", err.Error())
	}

	defer wg.Done()
}

// Runs python script that generates the toporgs.json file containing a table
// of the organizations that pay the most.
func CreateBestOrgsTable(wg *sync.WaitGroup) {
	logs.Info("Python", "Starting topOrg.py script.")

	cmd := exec.Command("../../use_cases/topOrg.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		logs.Errorf("Python", err.Error())
	}

	defer wg.Done()
}

// Runs python script that generates the toppos.json file containing a table of
// the positions that earn the most.
func CreateBestPosTable(wg *sync.WaitGroup) {
	logs.Info("Python", "Starting topPos.py script.")

	cmd := exec.Command("../../use_cases/topPos.py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logs.Errorf("Python", err.Error())
	}

	defer wg.Done()
}
