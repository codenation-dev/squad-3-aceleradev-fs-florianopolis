package main

import (
	"fmt"
	"net/http"
	"squad-3-aceleradev-fs-florianopolis/entities/logs"
	api "squad-3-aceleradev-fs-florianopolis/interfaces/API"
	"github.com/gorilla/handlers"
)

func main() {
	R := api.Initialize() 
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	logs.Errorf("Router ", fmt.Sprintf("Server dead %s", http.ListenAndServe(":8921", handlers.CORS(allowedHeaders,allowedMethods,allowedOrigins)(R.Router)).Error()))
}
