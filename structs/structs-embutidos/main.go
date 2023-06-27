package main

import "fmt"

func main() {

	type pessoa struct {
		nome  string
		idade int
	}

	type profissional struct {
		pessoa
		profissao string
		salario   int
	}

	pessoa1 := profissional{

		pessoa: pessoa{
			nome:  "arlindo",
			idade: 21,
		},
		profissao: "pizzaiolo",
		salario:   1000000,
	}

	pessoa2 := pessoa{
		nome:  "Eduardo",
		idade: 73,
	}

	fmt.Println(pessoa1)

	fmt.Println(pessoa2)
}
