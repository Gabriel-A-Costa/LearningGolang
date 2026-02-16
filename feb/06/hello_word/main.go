package main

import (
	"log"
	//"fmt"
	"net/http"
	"github.com/gabriel-a-costa/LearningGolang/feb/06/hello_world/handler"

)

func main() {	
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /hello", handler.HandleHelloWorld)
	mux.HandleFunc("POST /ping", handler.HandlePingPong)
	
	log.Println("Server runner in port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Erro ao iniciar servido: ", err)
	}
}
