package main

// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {

		x := i % 4
		fmt.Println(x)

	}
}
