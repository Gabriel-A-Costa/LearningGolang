package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloWorld(writer io.Writer) (int, error) {
	data := []byte("Hello World")
	n, err := writer.Write(data)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	http.HandleFunc(
		"/status",
		func(w http.ResponseWriter, r *http.Request) {

			w.Header().
				Set("Content-type", "application/json")

			HelloWorld(w)
		},
	)

	fmt.Printf("Escultado a porta :8080...")
	http.ListenAndServe(":8080", nil)
}
