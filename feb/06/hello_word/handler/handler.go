package handler

import "net/http"

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
}

func HandlePingPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}


