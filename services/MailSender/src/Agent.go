package mail

import ("squad-3-aceleradev-fs-florianopolis/entities/logs"
"net/smtp"
"fmt"
"squad-3-aceleradev-fs-florianopolis/interfaces/crud/emailenviado")

//SetHost set the host of smtpagent
func (S *SMTPAgent) SetHost(H SMTPHost) {
	logs.Info("Mail - Agent",fmt.Sprintf("Setting the host to %s",H.toString()))
	S.Host = H
}

//CreateSender create a sender structure
func (S *SMTPAgent) CreateSender(login string,pass string) {
	S.Senders = append(S.Senders,S.Host.createSender(login,pass))
}

//Send send a email
func (S *SMTPAgent) Send() {
	logs.Info("Mail - Agent","Tryingo to send mail")
	var auth = S.Senders[S.senderCtrl].Authme()
	for k,v := range S.Mail.Targets{
		msg := TemplateIt(S.Mail,k)
		err := smtp.SendMail(S.Host.toString(),auth,S.Senders[S.senderCtrl].login,[]string{v.Mail},msg)
		if (err!=nil){
			logs.Errorf("Mail - Agent",fmt.Sprintf("Error %s",err.Error()))
			panic(err)
		} else {
			err = emailenviado.Insert(S.Mail.ID,v.Mail)
			if err != nil {
				logs.Errorf("EmailInsert",err.Error())
			}
		}
	}
}

//PasswordSend send a email with password
func (S *SMTPAgent) PasswordSend() {
	logs.Info("Mail - Agent","Tryingo to send mail")
	var auth = S.Senders[S.senderCtrl].Authme()
	msg := TemplatePass(S.Pass)
	logs.Info("This is what you are sending",string(msg))
	err := smtp.SendMail(S.Host.toString(),auth,S.Senders[S.senderCtrl].login,[]string{S.Pass.Target.Mail},msg)
	if (err!=nil){
		logs.Errorf("Mail - Agent",fmt.Sprintf("Error %s",err.Error()))
		panic(err)
	} 
}

func (S *SMTPAgent) controlSender() {
	if (S.senderCtrl == len(S.Senders)){
		S.senderCtrl = 0
	} else {
		S.senderCtrl++
	}
}

func (h *SMTPHost) toString() string {
	return fmt.Sprintf("%s:%s",h.Host,h.Port)
} 

func (h *SMTPHost) createSender(login string,pass string) Sender {
	logs.Info("Mail - Agent",fmt.Sprintf("Creating a sender with login %s",login))
	return Sender{server:*h,login:login,pass:pass}
}

//Authme authenticate the sender
func (s *Sender) Authme() smtp.Auth {
	return smtp.PlainAuth("", s.login,s.pass, s.server.Host)
}

