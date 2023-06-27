package main

// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.

import "fmt"

func main() {

	num := [5]int{}

	num[0] = 1
	num[1] = 2
	num[2] = 10
	num[3] = 4

	for i, v := range num {

		fmt.Println(i, v)
	}

	fmt.Printf("%T", num)

}
