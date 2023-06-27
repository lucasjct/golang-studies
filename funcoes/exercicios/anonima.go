// You can edit this code!
// - Crie e utilize uma função anônima.
package main

import "fmt"

func main() {
	x := 5

	func(x int) {

		fmt.Println(x)

	}(x) //chamando a função passando x como argumento
}
