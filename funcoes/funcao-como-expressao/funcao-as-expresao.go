package main

import "fmt"

func main() {

	x := 2

	y := func(x int) int {
		//fmt.Println("o resultado de 25 multiplicado por", x, "é:", y)
		return x * 25
	}
	// na função como expressão, podemos atribuir um valor para uma variável a partir do returno de uma função.
	fmt.Println("o resultado de 25 multiplicado por", x, "é:", y(x))
}
