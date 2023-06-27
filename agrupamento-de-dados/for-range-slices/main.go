package main

import (
	"fmt"
)

// o for range é utilizado para fazer iterações no interior dos slices.
// o for range retorna sempre o índice e o valor do dados dentro do slice.

func main() {

	frutas := []string{"Banana", "Laranja", "Maçã"}

	frutas = append(frutas, "Melão")

	for indice, valor := range frutas {

		fmt.Println("No índice", indice, "temos a fruta:", valor)
	}

	//! exemplo com for range iterando dentro de um slice com tipos int

	num := []int{2, 4, 6, 8, 10}
	tot := 0

	for _, s := range num {
		tot += s

	}
	fmt.Println(tot)
}
