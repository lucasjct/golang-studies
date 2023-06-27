//    - Crie uma função que retorne um int
//    - Crie outra função que retorne um int e uma string
//    - Chame as duas funções
//    - Demonstre seus resultados

package main

import "fmt"

func main() {
	x := retornarString()
	x()
}

func retornarString() func() {

	return func() {

		fmt.Println("oi")

	}

}
