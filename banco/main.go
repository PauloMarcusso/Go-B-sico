package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

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

	// contaDaMalu := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Malu", CPF: "033033003003", Profissao: "Developer"},
	// 	NumeroAgencia: 123, NumeroConta: 123321}

	// titularPaulo := clientes.Titular{"Paulo H", "033033033033033", "Desevolvedor Go"}
	// contaPaulo := contas.ContaCorrente{Titular: titularPaulo, NumeroAgencia: 123, NumeroConta: 123123}
	// contaPaulo.Depositar(3000)

	// fmt.Println(contaDaMalu)
	// fmt.Println(contaPaulo)
	// fmt.Println("O saldo do", contaPaulo.Titular.Nome, "é de ", contaPaulo.ObterSaldo())

	contaDaJulia := contas.ContaPoupanca{}
	contaDaJulia.Depositar(1000)
	PagarBoleto(&contaDaJulia, 500)

	fmt.Println(contaDaJulia.ObterSaldo())
}
