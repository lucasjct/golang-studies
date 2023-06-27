package main

import "fmt"

func main() {

	sabores := []string{"abacaxi", "calabresa", "marguerita", "quatro queijos"}

	fmt.Println(sabores[0:2])

	//! exibir conteúdo do slice a partir do for retornar o índice.
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	//! deletar valor de uma slice

	sabor := append(sabores[:3], sabores[:3]...)
	fmt.Println(sabor)

}
