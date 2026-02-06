package operacao

import "fmt"

type Soma struct{}

func (s Soma) Calcular(a, b float64) float64 {
	return a + b
}

type Subtracao struct{}

func (s Subtracao) Calcular(a, b float64) float64 {
	return a - b
}

type Multiplicacao struct{}

func (m Multiplicacao) Calcular(a, b float64) float64 {
	return a * b
}

type Divisao struct{}

func (d Divisao) Calcular(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Erro: Divisão por zero!")
		return 0
	}
	return a / b
}
