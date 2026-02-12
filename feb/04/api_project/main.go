package main

import (
	"fmt"
	"net/http"

	"github.com/gabriel-a-costa/api_project/handler"
	"github.com/gabriel-a-costa/api_project/middleware"
)

func main() {
	mux := http.NewServeMux()

	loggedMux := middleware.LogginMiddleware(mux)

	mux.HandleFunc("GET /users", handler.ListUsers)
	mux.HandleFunc("GET /users/", handler.GetUserById)

	fmt.Println("Server Runner in port 8080...")
	http.ListenAndServe(":8080", loggedMux)
}
