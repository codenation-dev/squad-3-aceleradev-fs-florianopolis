package mail

import ("github.com/gorilla/mux"
"net/http"
"encoding/json"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"fmt")

func (S *SMTPAgent) getPost(w http.ResponseWriter,r *http.Request) {
	logs.SimpleRequest("MailSender_Listener_getPost","Received a request")
	decoder := json.NewDecoder(r.Body)
	var m Mailrequest 
	err := decoder.Decode(&m)
	if(err!=nil){
		logs.Errorf("MailSender_Listener_getPost",fmt.Sprintf(" Erro in decoding json: %s",err.Error()))
		panic(err)
	}
	S.Mail = m
	S.Send()
}

//Router return a Router
func Router(S *SMTPAgent) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/send",S.getPost).Methods("POST")
	return router
}

