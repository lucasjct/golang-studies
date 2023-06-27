// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	i := Inception() // atribuo a função principal para i
	i()              // evoco a função que está na variável 'i'.
}

// momento em que nomeia a função e define o tipo que ela retonar, no caso retorna uma func
func Inception() func() {
	return func() { // no retorno, deve ser declarada uma função anônima que será o retorno da função.

		fmt.Println("camadas dentro de camadas")

	}

}
