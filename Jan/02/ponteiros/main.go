package main

import "fmt"

func main() {
	idade := 30

	ponteiroIdade := &idade

	fmt.Println("Valor de idade:", idade)

	fmt.Println("Endereco de idade:", ponteiroIdade)

	fmt.Println("valor via ponteiro:", *ponteiroIdade)
}
