package main

import "fmt"

// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

func main() {

	hobbies := map[string][]string{
		"teixeira_lucas":   {"comer", "beber café"},
		"giaquinto_marina": {"conversar", "commer pipoca"},
		"Ana":              {"conversar", "Ficar com o Gael"},
	}

	for chave, valor := range hobbies {
		fmt.Println(chave)
		for i, v := range valor {
			fmt.Println("\t", i, v)

		}
	}

}
