package main

import ("net/http"
"squad-3-aceleradev-fs-florianopolis/interfaces/db"
"squad-3-aceleradev-fs-florianopolis/interfaces/Router"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"fmt")

func main() {
	F := database.FakeDB{User:"Leo",Pass:"Leo"}
	R := router.StartR(F)
	logs.Errorf("Router ",fmt.Sprintf("Server dead %s",http.ListenAndServe(":8921", R.Router).Error()))
}

/*
import ("squad-3-aceleradev-fs-florianopolis/interfaces/mail")

func main() {
	TgtList := mailApi.TargetList{}
	TgtList.Add("Felipe","psychelipe@gmail.com")
	var FakeINFO []string
	FakeINFO = append(FakeINFO,"bla")
	FakeINFO = append(FakeINFO,"bla2") 
	OtherInfo := "MsgInfoAndOthers"
	mailApi.Send(OtherInfo,TgtList,OtherInfo,FakeINFO,FakeINFO,OtherInfo)
}
*/