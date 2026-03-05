package main

import (
	//"context"
	"log"
	//"fmt"
	"net/http"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/handler"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/service"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/first_service/repository"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	kivik "github.com/go-kivik/kivik/v4"

)

func main() {	
	//ctx := context.background()

	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente couchdb: %s", err)
		return
	}

	db := client.DB("notebook")
	if err := db.Err(); err != nil {
		log.Fatalf("erro ao conectar ao db: %s", err)
		return
	}

	mux := http.NewServeMux()

	repo := repository.New(db)
	srv := service.NewService(repo)
	
	h := handler.New(srv)
	h.MountHandlers(mux)

	log.Println("server runner in port :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("erro ao iniciar servido: ", err)
	}
}
