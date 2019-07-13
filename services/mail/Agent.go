package mail

import ("squad-3-aceleradev-fs-florianopolis/entities/logs"
"net/smtp"
"fmt")

//SMTPAgent defines a agent
type SMTPAgent struct {
	Host SMTPHost
	Senders []Sender
	senderCtrl int
	Targets []Target
	Info []string
	Subject string
}

//SetHost set the host of smtpagent
func (S *SMTPAgent) SetHost(H SMTPHost) {
	logs.Info("Mail - Agent",fmt.Sprintf("Setting the host to %s",H.toString()))
	S.Host = H
}

//CreateSender create a sender structure
func (S *SMTPAgent) CreateSender(login string,pass string) {
	S.Senders = append(S.Senders,S.Host.createSender(login,pass))
}

//AddTarget add a target to SMTPAgent
func (S *SMTPAgent) AddTarget(Name string, Mail string) {
	logs.Info("Mail - Agent",fmt.Sprintf("Adding target to list with name %s and mail %s",Name,Mail))
	T := Target{Name:Name,Mail:Mail}
	S.Targets = append(S.Targets, T) 
}

//SetInfo set the message contents
func (S *SMTPAgent) SetInfo(Info []string) {
	logs.Info("Mail - Agent",fmt.Sprintf("Set the content of message"))
	S.Info = Info
}

//Send a email
func (S *SMTPAgent) Send() {
	logs.Info("Mail - Agent","Trying to send mail")
	var auth = S.Senders[S.senderCtrl].Authme()
	for _,v := range S.Targets{
		msg := TemplateIt(v.Name,S.Info)
		err := smtp.SendMail(S.Host.toString(),auth,S.Senders[S.senderCtrl].login,[]string{v.Mail},msg)
		if (err!=nil){
			logs.Errorf("Mail - Agent",fmt.Sprintf("Error %s",err.Error()))
			panic(err)
		}
	} 

}

func (S *SMTPAgent) controlSender() {
	if (S.senderCtrl == len(S.Senders)){
		S.senderCtrl = 0
	} else {
		S.senderCtrl++
	}
}

//SMTPHost references the host of smtp service
type SMTPHost struct {
	Host string 
	Port string
}  

func (h *SMTPHost) toString() string {
	return fmt.Sprintf("%s:%s",h.Host,h.Port)
} 

func (h *SMTPHost) createSender(login string,pass string) Sender {
	logs.Info("Mail - Agent",fmt.Sprintf("Creating a sender with login %s",login))
	return Sender{server:*h,login:login,pass:pass}
}

//Sender struct references to a sender
type Sender struct {
	server SMTPHost
	login string
	pass string
}

//Target struct refeto email to send
type Target struct {
	Name string
	Mail string
}

//Authme authenticate the sender
func (s *Sender) Authme() smtp.Auth {
	return smtp.PlainAuth("", s.login,s.pass, s.server.Host)
}

