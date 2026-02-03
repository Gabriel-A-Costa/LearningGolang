package saudacao

import "fmt"

func sayHi(nome string) string {
	return fmt.Sprintf("Bem vindo, %s\n", nome)
}

func BoasVindas(nome string) {
	fmt.Println(sayHi(nome))
}
