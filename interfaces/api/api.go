package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type listaClientes struct {
	Nome string `json:"nome"`
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/api/postclients", handleClients()).Methods("POST")
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleClients() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		list := csv.NewReader(r.Body)

		var structuredList []listaClientes

		for {
			line, err := list.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			cliente := listaClientes{
				Nome:    line[0],
			}

			structuredList = append(structuredList, cliente)
		}

		jason, err := json.Marshal(structuredList)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("clientlist.json", jason, 0644)
})
}
