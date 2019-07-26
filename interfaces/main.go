package main

import (
	"fmt"
	"net/http"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	api "squad-3-aceleradev-fs-florianopolis/interfaces/API"
)

func main() {
	R := api.Initialize()
	CreateNotify()
	logs.Errorf("Router ", fmt.Sprintf("Server dead %s", http.ListenAndServe(":8921", R.Router).Error()))
}
