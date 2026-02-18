package main

import (
	"context"
	"log"
	//"fmt"
	"net/http"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/hello_world/handler"
	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"

)

func main() {	
	//ctx := context.Background()

	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("Erro ao criar cliente CouchDB: %s", err)
	}

	db := client.DB("notebook")
	if err := db.Err(); err != nil {
		log.Fatalf("Erro ao conectar ao DB: %s", err)
	}

	mux := http.NewServeMux()
	
	h := handler.New(db)

	mux.HandleFunc("GET /hello", h.HandleHelloWorld)
	mux.HandleFunc("POST /ping", h.HandlePingPong)
	
	log.Println("Server runner in port :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Erro ao iniciar servido: ", err)
	}
}
