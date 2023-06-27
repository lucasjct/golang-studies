package main

import "fmt"

func main() {
	a := x()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := x()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

// exemplo de Closure,
// conseguimos acessar a variável que retornar na função interna independente do seu escopo.
func x() func() int {

	x := 0
	return func() int {
		x++
		return x
	}

}
