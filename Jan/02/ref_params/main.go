package main

import "fmt"

func alterarPonteiro(x *int) {
	*x = *x + 2
}

func main() {
	numero := 10
	fmt.Println("Definido como:", numero)
	alterarPonteiro(&numero)
	fmt.Println("Fora da função:", numero)
}
