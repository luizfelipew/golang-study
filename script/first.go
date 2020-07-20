package main

import (
	"fmt"
)

func main() {
	nome := "Luiz"
	versao := 1.1
	fmt.Println("Ola Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	// fmt.Println("O tipo da variavel eh:", reflect.TypeOf(versao))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("O endereco da variavel comando eh", &comando)
	fmt.Println("O comando escolhido foi:", comando)
}
