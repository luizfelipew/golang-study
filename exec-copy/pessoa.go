package main

import "fmt"

type Pessoa struct {
	nome, sobrenome string
}

func (p Pessoa) ExibirNomeCompleto() string {
	p.sobrenome = "Campos"
	nomeCompleto := p.nome + " " + p.sobrenome
	return nomeCompleto
}

func main() {
	p1 := Pessoa{"Luiz", "Wendt"}
	fmt.Println(p1.ExibirNomeCompleto())
	fmt.Println(p1.nome, p1.sobrenome)
}
