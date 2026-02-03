package main

import "fmt"

func main() {
	numeros := []int{10, 20, 30}
	for i, numero := range numeros {
		fmt.Printf("Index %d: %d, %v \n", i, numero, &numero)
	}
}
