package main

import "fmt"

// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

func main() {

	data := [][]string{
		{"Nome", "Sobre Nome", "Hobby favorito"},
		{"Lucas", "Teixeira", "Futebol"},
		{"Marina", "Giaquinto", "Conversar"},
		{"Paulo", "Teixeira", "Musica"},
	}

	for _, i := range data {
		fmt.Println(i)
	}

}
