package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getUserHandle(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	w.Header().
		Set("Content-type", "application/json")

	json.NewEncoder(w).
		Encode(
			map[string]string{"id": id, "status": "ok"},
		)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users/", getUserHandle)
	fmt.Println("Running in port 8080")
	http.ListenAndServe(":8080", mux)
}
