package utils

import (
funcpublico "squad-3-aceleradev-fs-florianopolis/interfaces/crud/funcpublico"
"squad-3-aceleradev-fs-florianopolis/entities"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"net/http"
"bytes"
"encoding/json"
"math/rand"
"time"
)

type MailType entity.Target

func getTopClientsName() (nomes []string) {
	nomes, erro := funcpublico.GetClienteFuncPublico()
	if erro != nil {
		logs.Errorf("getNames", erro.Error())
	}
	return nomes
}

//Funcionários públicos com maiores salários (acima de 20k, top 10)
func getTopIncomes() (nomes []string) {
	nomes, erro := funcpublico.GetTop10Incomes()
	if erro != nil {
		logs.Errorf("getNames", erro.Error())
	}
	return nomes
}

//RequestCreator Utils Return a request
func RequestCreator(Targets []entity.Target, ID int) entity.Mailrequest {
	var request entity.Mailrequest
	request.Subject = "Novas oportunidades UATI"
	request.Targets = Targets
	request.Message = "Aqui temos as oportunidades da semana..."
	request.TopNames = getTopClientsName()
	request.Names = getTopIncomes()
	request.Link = "linkdobotao" //ALTERAR
	request.ID = ID
	return request
}

//SendMailRequest Send a Mail Request
func SendMailRequest(r entity.Mailrequest) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)
	_,es := http.Post("http://127.0.0.1:8225/send", "application/json; charset=utf-8", b)
	if (es!=nil){
		logs.Errorf("MailConsumer - SendMailNotify",es.Error())
	}
}
 
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  
func GeneratePassword() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))  
	b := make([]byte, 8)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
		return string(b)
}
  
func ConvertDateTimeSQL(date string) time.Time {
	var d time.Time 
	d, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return d
	}
	return d
}