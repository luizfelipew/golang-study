package main

import (
	"fmt"

	"./contas"
)

func main() {
	contaDaPati := contas.ContaCorrente{Titular: "Pati", Saldo: 300}
	contaDoLuiz := contas.ContaCorrente{Titular: "Luiz", Saldo: 100}

	status := contaDoLuiz.Transferir(200, &contaDaPati)

	fmt.Println(status)
	fmt.Println(contaDaPati)
	fmt.Println(contaDoLuiz)
}
