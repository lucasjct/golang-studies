package main

import "fmt"

func main() {

	// com defer, a execução da função é adiada e executada no final do programa.
	// podemos utilizar defer para garantir que se encerre uma execuçaoao final do programa.
	defer fmt.Println("1 - executar com defer")
	fmt.Println("2 - executar sem defer")
}
