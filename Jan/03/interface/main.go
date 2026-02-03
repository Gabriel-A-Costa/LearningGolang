package main

import "fmt"

type Carro struct {
	Modelo string
}

func (c Carro) Acelerar() {
	fmt.Printf("O carro %s esta acelerando\n", c.Modelo)
}

type Bicicleta struct {
	Tipo string
}

func (b Bicicleta) Acelerar() {
	fmt.Printf("A Bicicleta %s esta acelerando\n", b.Tipo)
}

type Veiculo interface {
	Acelerar()
}

func TestarVeiculo(v Veiculo) {
	v.Acelerar()
}

func main() {
	carro := Carro{Modelo: "Ferrari"}
	bike := Bicicleta{Tipo: "Toshiba"}

	TestarVeiculo(carro)
	TestarVeiculo(bike)
}
