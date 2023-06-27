package main

import "fmt"

func main() {

	// com o make, podemos escalar a capacidade do array.
	// podemos ver abaixo que a capacidade do array dobra, caso a quantidade de itens atinja o limite max.

	slice := make([]int, 5, 10)

	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)

	fmt.Println("Tamanho do slice:", len(slice), "\ncapacidade do slice", cap(slice), "\n elementos do slice", slice)

}
