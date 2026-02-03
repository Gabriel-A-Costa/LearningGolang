package main

import (
	"fmt"
	"time"
)

func bornDate(age int) {
	year := time.Now().Year()
	born := year - age
	fmt.Printf("Voce tem %d, e nasceu em %d \n", age, born)
}

func main() {
	var age int
	for {
		fmt.Println("Digite sua idade:")
		_, err := fmt.Scanln(&age)

		if err != nil {
			fmt.Println("Entrada invalida, pfv digite um numero")
			continue
		}

		if age < 18 {
			fmt.Println("Voce precisa ser maior de idade")
			continue
		}
		break
	}

	bornDate(age)
}
