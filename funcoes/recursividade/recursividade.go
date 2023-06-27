package main

import "fmt"

func main() {

	fmt.Println(fatorial(6))
	fmt.Println(loopFatorial(6))

}

// função recursiva - dentro da definição da função é chamada a própria função. Por isso é recursiva.
// Efeito drosta e fractal é um modelo para pensar funções recursivas.
func fatorial(x int) int {

	if x == 1 {
		return 1
	}
	return fatorial(x) * fatorial(x-1)
}

// fatorial feita com loop, loops consomem menos memória do que recursividade, devemos privilegiar loops.
func loopFatorial(x int) int {
	t := 1
	for x > 1 {

		t *= x
		x--

	}
	return t
}
