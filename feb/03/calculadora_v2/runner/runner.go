package runner

import (
	"fmt"

	"github.com/gabriel-a-costa/calculadoraV2/operacao"
)

type Operacao interface {
	Calcular(a, b float64) float64
}

type Runner struct {
	Operacoes map[string]Operacao
}

func (r *Runner) Excultar() {
	var a, b float64
	var operacao string
	fmt.Println("Digite o primeiro numero:")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo numero:")
	fmt.Scanln(&b)
	fmt.Println("Escolha a operação (+, -, *, /):")
	fmt.Scanln(&operacao)

	op, existe := r.Operacoes[operacao]
	if !existe {
		fmt.Println("Operação Inválida!")
		return
	}

	result := op.Calcular(a, b)
	fmt.Printf("Resultado: %.2f\n", result)
}

func NewRunner() *Runner {
	return &Runner{
		Operacoes: map[string]Operacao{
			"+": operacao.Soma{},
			"-": operacao.Subtracao{},
			"*": operacao.Multiplicacao{},
			"/": operacao.Divisao{},
		},
	}
}
