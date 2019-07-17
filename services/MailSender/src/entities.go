package mail

//Mailrequest Type for sending a mail
type Mailrequest struct {
	Subject string `json:"Subject"`
	Targets []Target `json:"Targets"`
	Message string `json:"Message"`
	TopNames []string `json:"TopNames,omitempty"`
	Names []string `json:"Names"`
	Link string `json:"Link"`
}

//Target struct refer email to send
type Target struct {
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

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
}

//Sender struct references to a sender
type Sender struct {
	server SMTPHost
	login string
	pass string
}