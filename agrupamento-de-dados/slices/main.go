package main

import "fmt"

func main() {
	// No slice não temos um número fixo expecificado
	// Sempre que um slice é redimensionado, um novo array é criado por debaixo dos panos.

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	array2 := append(slice, 6)

	fmt.Println(array2)
	fmt.Println(slice[2])
}
