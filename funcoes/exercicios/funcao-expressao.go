// - Atribua uma função a uma variável.
// - Chame essa função.
package main

import "fmt"

func main() {
	y := "Oi"
	x := func(x string) string {

		return x

	}
	fmt.Println(x(y))
}
