package main

import "fmt"

// Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

func main() {

	anonascimento := 1989
	anoatual := 2023

	for anonascimento <= anoatual {

		// o contador não pode ser posicionado nesta linha porque ele inicializaria com o ano de 1990 e naõ 1989
		// isso porque o programa é lido de cima para baixo.
		fmt.Println(anonascimento)
		anonascimento++

	}

}
