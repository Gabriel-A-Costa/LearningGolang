package main

import (
	//"context"
	"log"
	//"fmt"
	"net/http"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/hello_world/handler"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	kivik "github.com/go-kivik/kivik/v4"

)

func main() {	
	//ctx := context.Background()

	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("Erro ao criar cliente CouchDB: %s", err)
		return
	}

	db := client.DB("notebook")
	if err := db.Err(); err != nil {
		log.Fatalf("Erro ao conectar ao DB: %s", err)
		return
	}

	mux := http.NewServeMux()
	
	h := handler.New(db)

	mux.HandleFunc("GET /heath", h.Heath)
	mux.HandleFunc("POST /notebook", h.Create)
	mux.HandleFunc("GET /notebook/{id}", h.Get)
	mux.HandleFunc("PUT /notebook/{id}", h.Update)
	mux.HandleFunc("DELETE /notebook/{id}/{rev}", h.Delete)
	mux.HandleFunc("GET /notebook/list", h.List)

	log.Println("Server runner in port :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Erro ao iniciar servido: ", err)
	}
}
