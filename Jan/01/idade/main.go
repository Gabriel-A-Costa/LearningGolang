package main

import "fmt"

func sumFunc(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var num1, num2 int
	fmt.Println("Digite um numero")
	fmt.Scanln(&num1)

	fmt.Println("Digite outro numero")
	fmt.Scanln(&num2)

	sum := sumFunc(num1, num2)

	fmt.Printf("A soma de %d e %d é igual a %d \n", num1, num2, sum)
}
