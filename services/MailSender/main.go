package main

import("squad-3-aceleradev-fs-florianopolis/entities/logs"
"squad-3-aceleradev-fs-florianopolis/services/MailSender/src"
"net/http"
"fmt")

func main(){
	Host := mail.SMTPHost{Host:"smtp.gmail.com",Port:"587"}
	Agent := mail.SMTPAgent{}
	Agent.SetHost(Host)
	Agent.LoadSenders()
	logs.Info("Mail - # ","Initializing mail service test")
	theRouter := mail.Router(&Agent)
	logs.Errorf("Mail - # ",fmt.Sprintf("Server dead %s",http.ListenAndServe(":8921", theRouter).Error()))
}

    