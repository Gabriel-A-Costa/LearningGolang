package main

import "fmt"

func idadeFunc(age int) {
	if age >= 18 {
		fmt.Println("Voce é maior de idade")
	} else {
		fmt.Println("voce é menor de idade")
	}
}

func main() {
	var age int
	fmt.Println("Qual sua idade")
	fmt.Scanln(&age)

	idadeFunc(age)
}
