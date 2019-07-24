package mailApi

import ("squad-3-aceleradev-fs-florianopolis/entities/logs"
"net/http"
"encoding/json"
"bytes")

type request struct {
	Subject string `json:"Subject"`
	Targets []target `json:"Targets"`
	Message string `json:"Message"`
	TopNames []string `json:"TopNames,omitempty"`
	Names []string `json:"Names"`
	Link string `json:"Link"`
}

//Target Type
type target struct {
	Name string `json:"Name"`
	Mail string `json:"Mail"`
}

//TargetList Generate TargetList
type TargetList struct {
	Targets []target
}

func (T TargetList) list() []target {
	return T.Targets
}

//Add A Target To List
func (T *TargetList) Add(Name string,Email string) {
	T.Targets = append(T.Targets, target{Name:Name,Mail:Email})
}

//Send a mail
func Send(Subject string,targets TargetList,Message string,TopNames []string,Names []string,Link string) {
	r := request{
		Subject: Subject,
		Targets: targets.Targets,
		Message:Message,
		TopNames:TopNames,
		Names:Names,
		Link:Link}
	J,err := json.Marshal(r)
	if (err!=nil) {
		logs.Errorf("mailApi","MarshalError"+err.Error())
		return
	}
	url := "http://localhost:8921/send"
	req,_ := http.NewRequest("POST",url,bytes.NewBuffer(J)) 
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_,err = client.Do(req)
	if (err!=nil) {
		logs.Errorf("mailApi","RequestError"+err.Error())
	}
}