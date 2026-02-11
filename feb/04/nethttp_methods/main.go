package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/users",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
					return
				}

				var user User
				if err := json.Unmarshal(body, &user); err != nil {
					http.Error(w, "Dados inválidos", http.StatusBadRequest)
					return
				}
				fmt.Fprintf(w, "Usuário %s com %d anos recebido!\n", user.Name, user.Age)
			} else {
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		})
	fmt.Println("Escultando porta: 8080...")
	http.ListenAndServe(":8080", nil)
}
