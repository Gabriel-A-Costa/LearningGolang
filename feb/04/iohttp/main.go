package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/status",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().
				Set("Content-Type", "application/json")

			json.NewEncoder(w).
				Encode(map[string]bool{"status": true})
		},
	)
	fmt.Println("Servindo na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
