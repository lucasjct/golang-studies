package main

import "fmt"

// a função append faz parte do pacote builin da linguagem

func main() {

	elem := []string{}

	e := "olá"
	e2 := "Tudo bem?"

	// os parâmetros são variáticos.
	elemnOne := append(elem, e, e2, "Como vão as coisas por aí?")

	fmt.Println(elemnOne)

	//! append com slices inteiros

	recebe := []int{}

	// tipo slice de int
	elementos := []int{2, 4, 6, 8, 10}

	// os "..." remetem-se ao conteúdo da fatia do slices presente na variável elementos.
	tot := append(recebe, elementos...)

	fmt.Println(tot)

}
