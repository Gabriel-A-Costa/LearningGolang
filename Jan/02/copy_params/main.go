package main

import "fmt"

func alterarCopia(x int) int {
	fmt.Println("Recebido como:", x)
	x = x + 2
	fmt.Println("Atualizado para:", x)
	return x
}

func main() {
	numero := 10
	fmt.Println("Definido como:", numero)
	y := alterarCopia(numero)
	fmt.Println("Fora da função:", y)
}
