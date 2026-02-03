package main

import (
	"errors"
	"fmt"
)

type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

/* ---------------  METODOS DA CALCULADORA  ------------------*/
func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, errors.New("Divisão por zero não é permitida.")
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

type Runner struct {
	calculadora *Calculadora

	valor1    float64
	valor2    float64
	resultado float64

	operacao string
}

/* ---------------  METODOS DO RUNNER  ------------------*/
func (r *Runner) SolicitarValor() error {
	fmt.Println("Digite o primeiro numero:")
	if _, err := fmt.Scanln(&r.calculadora.Operando1); err != nil {
		return errors.New("entrada invalida para o primeiro numero")
	}

	fmt.Println("Digite o segundo numero:")
	if _, err := fmt.Scanln(&r.calculadora.Operando2); err != nil {
		return errors.New("entrada invalida para o segundo numero")
	}

	return nil
}

func (r *Runner) solicatarOperacao() error {
	fmt.Println("Escolha a operacao ( + - / *)")
	if _, err := fmt.Scanln(&r.operacao); err != nil {
		return errors.New("entrada invalida para operação")
	}

	switch r.operacao {
	case "+", "-", "/", "*":
		return nil
	default:
		return errors.New("operação invalida.")
	}
}

func (r *Runner) execultarOperacao() {
	switch r.operacao {
	case "+":
		fmt.Println("Resultado", r.calculadora.Somar())
	case "-":
		fmt.Println("Resultado", r.calculadora.Subtrair())
	case "*":
		fmt.Println("Resultado", r.calculadora.Multiplicar())
	case "/":
		result, err := r.calculadora.Dividir()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Resultado", result)
		}
	}
}

func (r *Runner) Execulte() {
	for {
		if err := r.SolicitarValor(); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		err := r.solicatarOperacao()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		r.execultarOperacao()
	}
}

/* ------ fUNÇÃO PARA ADD STRUCT CALCULADORA NO RUNNER  -----*/
func NewRunner(c *Calculadora) *Runner {
	return &Runner{calculadora: c}
}

func main() {
	calculadora := &Calculadora{}
	runner := NewRunner(calculadora)
	runner.Execulte()
}
