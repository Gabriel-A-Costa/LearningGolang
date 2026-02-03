package main

import "fmt"

func main() {
	var name string
	fmt.Println("Qual é seu nome:")
	fmt.Scan(&name)

	fmt.Printf("Ola %s, seja bem vindo!\n", name)
}
