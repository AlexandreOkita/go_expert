package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func Desativar(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	okita := Cliente{
		Nome:  "Okita",
		Idade: 24,
		Ativo: true,
	}
	Desativar(okita)

	fmt.Println(okita)
	fmt.Printf("Nome: %s Idade: %d Ativo %d", okita.Nome, okita.Idade, okita.Ativo)
}
