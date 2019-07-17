package mail

import ("io/ioutil"
"strings"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"fmt")

//LoadSenders Load Credentials from file
func (A *SMTPAgent) LoadSenders() {
	
	F,err := ioutil.ReadFile(".cfg/credentials")	
	
	if (err!=nil) {
		logs.Errorf("Mail_config_LoadCredentials",err.Error())
		panic(err)
	}

	Lines := strings.Split(string(F),"\n")

	for _,V := range Lines {
		L := strings.Split(V,",")
		A.CreateSender(L[0],L[1])
		logs.Info("Mail_config_LoadCredentials",fmt.Sprintf("Loaded Sender %s",L[0]))
	}

}