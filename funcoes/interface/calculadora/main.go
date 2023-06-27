package main

import "fmt"

type calc interface {
	soma() int
	divisao() int
	multiplicacao() int
	subtracao() int
}

type operators struct {
	num1, num2 int
}

func (o operators) soma() int {
	return o.num1 + o.num2

}

func (o operators) divisao() (string int) {

	return o.num1 / o.num2
}

func (o operators) multiplicacao() int {
	return o.num1 * o.num2
}

func (o operators) subtracao() int {
	return o.num1 - o.num2
}

func calcular(c calc) {
	fmt.Println(c.soma())
	fmt.Println(c.divisao())
	fmt.Println(c.multiplicacao())
	fmt.Println(c.subtracao())
}
func main() {

	o := operators{20, 2}

	calcular(o)

}
