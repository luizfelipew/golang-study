package main

import "fmt"

func Somando(numeros ...int) int {
	resultadoSoma := 0
	for _, numero := range numeros {
		resultadoSoma += numero
	}
	return resultadoSoma
}

func main() {
	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
}
