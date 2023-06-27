//- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
//- Passe um valor do tipo slice of int como argumento para a função;
//- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
//- Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 6}
	fmt.Println(somaInt(1, 1, 2, 1, 1))
	fmt.Println(somaSliceInt(sliceInt))
}

func somaInt(num ...int) int {

	total := 0
	for _, i := range num {

		total += i

	}
	return total

}

func somaSliceInt(num []int) int {
	soma := 0
	for _, i := range num {

		soma += i

	}
	return soma
}
