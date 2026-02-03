package main

import "fmt"

type Pessoa struct {
	Nome  string
	Maior bool
	Idade int
}

// Metodos
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Ola, eu sou %s e tenho %d anos", p.Nome, p.Idade)
}

func (p *Pessoa) Envelhecer() {
	p.Idade++
}

func main() {
	aluno := Pessoa{"Gabriel", true, 25}

	fmt.Println("Aluno.Nome:", aluno.Nome)
	fmt.Println("Aluno.Maior:", aluno.Maior)
	fmt.Println("Aluno.Idade:", aluno.Idade)
	fmt.Println("Struct Pessoa:", aluno)

	fmt.Println("Aluno.Nome Ponteiro:", &aluno.Nome)
	fmt.Println("Aluno.Maior Ponteiro:", &aluno.Maior)
	fmt.Println("Aluno.Idade Ponteiro:", &aluno.Idade)

	fmt.Println("Apresentaçao:", aluno.Apresentar())
	aluno.Envelhecer()
	fmt.Println("Aluno.Idade:", aluno.Idade)
}
