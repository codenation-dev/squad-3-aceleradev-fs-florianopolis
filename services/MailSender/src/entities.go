package mail

import ("squad-3-aceleradev-fs-florianopolis/entities")

type PasswordRequest entity.PasswordRequest
type Mailrequest entity.Mailrequest

//SMTPHost references the host of smtp service
type SMTPHost struct {
	Host string 
	Port string
}  

//SMTPAgent defines a agent
type SMTPAgent struct {
	Host SMTPHost
	Senders []Sender
	senderCtrl int
	Mail Mailrequest
	Pass PasswordRequest
}

//Sender struct references to a sender
type Sender struct {
	server SMTPHost
	login string
	pass string
}
