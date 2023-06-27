// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// métodos é onde descobrimos como tipos são importantes em go
// métodos é uma função anexada a um tipo.
// o método adciona funcionalidade a um tipo

type pessoa struct {
	nome  string
	idade int
}

// a função ola recebeu p que é do tipo pessoa.
func (p pessoa) ola() {

	fmt.Println(p.nome, "diz olá")

}

func main() {
	m := pessoa{"Lucas", 33}
	m.ola()

}
