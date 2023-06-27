package main

import "fmt"

// Exemplo de struct anônimo. ele pode ser usada uma única vez.
func main() {

	x := struct {
		nome  string
		idade int
	}{
		"Mario",
		40,
	}
	fmt.Println(x.nome)
}

// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

func exec() {

	classificacao := struct {
		n int
		d string
	}{
		1,
		"doius",
	}

}
