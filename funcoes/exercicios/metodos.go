// You can edit this code!
// Click here and start typing.
package main

// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) NomeIdade() {

	fmt.Println("Sou", p.nome, p.sobrenome, "tenho", p.idade, "anos")
}

func main() {

	p := Pessoa{"Lucas", "Teixeira", 20}

	p.NomeIdade()
}
