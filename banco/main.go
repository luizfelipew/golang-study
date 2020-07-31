package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoLuiz := ContaCorrente{titular: "Luiz Felipe", numeroAgencia: 589,
		numeroConta: 123456, saldo: 125.50}

	contaDaPati := ContaCorrente{"Pati", 1000, 100000, 300.50}

	fmt.Println(contaDoLuiz)
	fmt.Println(contaDaPati)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
}
