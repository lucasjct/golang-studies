package main

import "fmt"

func main() {
	x := 20

	// função anônima não recebe um nome e é para ser 'descartável', sem
	// pretensão de usá-la em outros trechos do código.
	func(x int) {

		fmt.Println("o resultado desta função anônima é:", x*10)
	}(x)

	func() int {

		return 10
	}()
}
