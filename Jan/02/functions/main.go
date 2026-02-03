package main

import "fmt"

func dobrar(numero int) (int, error) {
	if numero < 0 {
		return 0, fmt.Errorf("Numero negativo %d", numero)
	}

	return numero * 2, nil
}

func main() {
	resultado, err := dobrar(-10)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("O resultado é:", resultado)
}
