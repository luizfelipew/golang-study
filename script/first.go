package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Näo conheço esse comando")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Luiz"
	versao := 1.1
	fmt.Println("Ola Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	// fmt.Println("O tipo da variavel eh:", reflect.TypeOf(versao))
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {

	var comandoLido int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)

	fmt.Println("O endereco da variavel comando eh", &comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "http://www.alura.com.br"
	http.Get(site)
}
