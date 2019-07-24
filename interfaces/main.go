package main

import ("net/http"
"squad-3-aceleradev-fs-florianopolis/interfaces/API"
"squad-3-aceleradev-fs-florianopolis/entities/logs"
"fmt")

func main() {
	R := api.Initialize()
	logs.Errorf("Router ",fmt.Sprintf("Server dead %s",http.ListenAndServe(":8921", R.Router).Error()))
}