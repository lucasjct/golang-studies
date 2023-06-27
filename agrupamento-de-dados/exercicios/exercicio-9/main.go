package main

import "fmt"

// - Utilizando o exercício anterior, adcione uma entrada do map e demonstre o map inteiro utilizando range.

func main() {

	hobbies := map[string][]string{
		"teixeira_lucas":   {"comer", "beber café"},
		"giaquinto_marina": {"conversar", "commer pipoca"},
		"Ana":              {"conversar", "Ficar com o Gael"},
	}

	hobbies["teixeira_paulo"] = []string{"ouvir música"}

	for chave, valor := range hobbies {
		fmt.Println(chave)
		for i, v := range valor {
			fmt.Println("\t", i, v)

		}
	}

}
