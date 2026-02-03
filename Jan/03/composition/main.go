package main

import "fmt"

type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s esta acelerando\n", c.Modelo)
}

func (c Carro) Freiar() {
	fmt.Printf("O carro %s esta freiando\n", c.Modelo)
}

type CarroEletrico struct {
	Carro
	BateriaCarga int
}

func (ce CarroEletrico) CarregarBateria() {
	fmt.Printf("O carro esta carregando... %d%%\n", ce.BateriaCarga)
}

type Veiculo interface {
	Acelerar()
	Freiar()
	CarregarBateria()
}

func TestarVeiculo(v Veiculo) {
	v.Acelerar()
	v.Freiar()
	v.CarregarBateria()
}

func main() {
	ce := CarroEletrico{
		Carro:        Carro{Modelo: "Tesla"},
		BateriaCarga: 80,
	}

	TestarVeiculo(ce)
}
