package main

import (
	"fmt"

	"./contas"
)

func main() {
	// contaDaPati := contas.ContaCorrente{Titular: "Pati", Saldo: 300}
	// contaDoLuiz := contas.ContaCorrente{Titular: "Luiz", Saldo: 100}

	// status := contaDoLuiz.Transferir(200, &contaDaPati)

	// fmt.Println(status)

	// contaDaPati := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Pati",
	// 	CPF:       "123.111.123.12",
	// 	Profissao: "Desenvolvedor GO"},
	// 	NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	// clientePati := clientes.Titular{"Pati", "123.111.123.12", "Desenvolvedor GO"}
	// contaDaPati := contas.ContaCorrente{clientePati, 123, 123456, 100}

	// fmt.Println(contaDaPati)

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())

}
