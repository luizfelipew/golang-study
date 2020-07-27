package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)

	exibeNomes()
	// exibeIntroducao()

	for {
		// exibeMenu()
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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!!!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Luiz", "Felipe", "Wendt"}

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
