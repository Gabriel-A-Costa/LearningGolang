package main

import "fmt"

type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, fmt.Errorf("Error: Divisão por zero")
	}
	return c.Operando1 / c.Operando2, nil
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operando1 * c.Operando2
}

func (c Calculadora) Somar() float64 {
	return c.Operando1 + c.Operando2
}

func (c Calculadora) Subtrair() float64 {
	return c.Operando1 - c.Operando2
}

func main() {
	var op1, op2 float64
	var operacao string

	fmt.Println("Digite o primeiro numero:")
	fmt.Scanln(&op1)
	fmt.Println("Digite o segundo numero:")
	fmt.Scanln(&op2)

	calc := Calculadora{op1, op2}

	fmt.Println("Escola uma operação: ( +, -, /, * )")
	fmt.Scanln(&operacao)

	switch operacao {
	case "+":
		fmt.Printf("A soma de %f e %f é: %f", op1, op2, calc.Somar())

	case "-":
		fmt.Printf("A subtração de %f e %f é: %f", op1, op2, calc.Subtrair())

	case "*":
		fmt.Printf("A multiplicação de %f e %f é: %f", op1, op2, calc.Multiplicar())

	case "/":
		val, err := calc.Dividir()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("A divisão de %f e %f é: %f", op1, op2, val)
	}

}
