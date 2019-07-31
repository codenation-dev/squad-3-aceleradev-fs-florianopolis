package mail

import ("fmt"
"html/template"
"bytes"
"squad-3-aceleradev-fs-florianopolis/entities/logs")

type templateStruct struct {
	TargetName string
	Names []string
	TopNames []string
	Link string
	Message string
}

func toTemplate(mail Mailrequest,index int) templateStruct {
	tpl := templateStruct{TargetName: mail.Targets[index].Name,
		Names: mail.Names,
		TopNames: mail.TopNames,
		Link: mail.Link,
		Message: mail.Message}
	return tpl
}

//TemplateIt Create a HTML and CSS
func TemplateIt(mail Mailrequest,index int) []byte{
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n "
	Header := fmt.Sprintf("To:%s\r\nSubject:%s\r\n",mail.Targets[index].Mail,mail.Subject)
	var tpl bytes.Buffer
	tmpl,err := template.ParseFiles("templates/templateone.html")
	if(err!=nil){
		logs.Errorf("MailSender_Templater_TemplateIt ",err.Error())
		panic(err)
	}
	tmpl.Execute(&tpl,toTemplate(mail,index))
	return []byte(Header+mime+tpl.String())
}

//TemplatePass Create a template html with css to password
func TemplatePass(Pass PasswordRequest) []byte {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n "
	Header := fmt.Sprintf("To:%s\r\nSubject:%s\r\n",Pass.Target.Mail,Pass.Subject)
	var tpl bytes.Buffer
	tmpl,err := template.ParseFiles("templates/templatepass.html")
	if(err!=nil){
		logs.Errorf("MailSender_Templater_TemplatePass ",err.Error())
		panic(err)
	}
	tmpl.Execute(&tpl,Pass)
	return []byte(Header+mime+tpl.String())
}