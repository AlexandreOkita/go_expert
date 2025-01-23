package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func somaEndereco(a, b *int) *int {
	*a = 50
	c := *a + *b
	return &c
}

type ContaBancaria interface {
	simularAcrescimo(valor int) int
	depositar(valor int)
	checarSaldo()
}

type Conta struct {
	saldo int
	nome string
}

func NewConta(saldo int, nome string) ContaBancaria {
	fmt.Printf("Conta do %v criada!\n", nome)
	return &Conta{saldo: saldo, nome: nome}
}

func (c *Conta) simularAcrescimo(valor int) int {
	c.saldo += valor
	fmt.Printf("Valor da simulação: %v\n", c.saldo + valor)
	return c.saldo + valor
}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
}

func (c *Conta) checarSaldo() {
	fmt.Printf("Saldo da conta: %v\n", c.saldo)
}

type ContaDoNubank struct {}

func (cn ContaDoNubank) simularAcrescimo(valor int) int {
	return 0
}
func (cn *ContaDoNubank) depositar(valor int) {}
func (cn ContaDoNubank) checarSaldo() {}
func NewContaDoNubank() ContaBancaria {
	return &ContaDoNubank{}
}

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	// de-referencing
	println(*b)
	*b = 30
	println(a)

	println("--------------")

	minhaVar1 := 10
	minhaVar2 := 20
	println(*somaEndereco(&minhaVar1, &minhaVar2))
	println(minhaVar1)

	println("--------------")

	conta := Conta{
		saldo: 100,
		nome: "Okita",
	}
	fmt.Printf("Conta do %v criada!\n", conta.nome)

	conta.simularAcrescimo(100)
	fmt.Printf("Saldo da conta: %v\n", conta.saldo)

	conta.depositar(100)
	fmt.Printf("Saldo da conta: %v\n", conta.saldo)

	conta2 := NewConta(0, "Jao")
	conta2.simularAcrescimo(100)
	conta2.checarSaldo()
	conta2.depositar(100)
	conta2.checarSaldo()
}