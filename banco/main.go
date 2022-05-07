package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	// contaDoPaulo := contas.ContaCorrente{Titular: "Paulo", NumeroAgencia: 111, NumeroConta: 222, Saldo: 1000.0}

	// valorDoSaque := 300
	// fmt.Println("Saldo antes do saque:", contaDoPaulo.Saldo)
	// contaDoPaulo.Sacar(float64(valorDoSaque))
	// fmt.Println("Saldo depois do saque:", contaDoPaulo.Saldo)
	// fmt.Println(contaDoPaulo)

	// contaDaPamela := contas.ContaCorrente{}
	// contaDaPamela.Saldo = 1000

	// fmt.Println(contaDaPamela)
	// contaDaPamela.Depositar(500)
	// fmt.Println(contaDaPamela)

	// contaDaLili := contas.ContaCorrente{Titular: "Lili", Saldo: 1000}
	// contaDaMaya := contas.ContaCorrente{Titular: "Maya", Saldo: 500}

	// contaDaLili.Transferir(430, &contaDaMaya)

	// fmt.Println(contaDaLili)
	// fmt.Println(contaDaMaya)

	contaDaMalu := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Paulo", CPF: "033033003003", Profissao: "Developer"},
		NumeroAgencia: 123, NumeroConta: 123321, Saldo: 5000.0}

	fmt.Println(contaDaMalu)
}
