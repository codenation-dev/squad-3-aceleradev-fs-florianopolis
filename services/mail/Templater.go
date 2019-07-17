package mail

import ("fmt"
"html/template"
"bytes"
"squad-3-aceleradev-fs-florianopolis/entities/logs")

type templateStruct struct {
	TargetName string
	Names []string
	TopNames []string
}

func toTemplate(mail Mailrequest,index int) templateStruct {
	tpl := templateStruct{TargetName: mail.Targets[index].Name,
		Names: mail.Names,
		TopNames: mail.TopNames}
	return tpl
}

//TemplateIt Create a HTML and CSS
func TemplateIt(mail Mailrequest,index int) []byte{
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n "
	Header := fmt.Sprintf("Subject:%s\r\n",mail.Subject)
	var tpl bytes.Buffer
	tmpl := template.Must(template.ParseFiles("mail/templateone.html"))
	tmpl.Execute(&tpl,toTemplate(mail,index))
	logs.Info("Template: ",tpl.String())
	return []byte(Header+mime+tpl.String())
}