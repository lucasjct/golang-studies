package main

import (
	"fmt"
	"math"
)

// As interfaces são compostas por um conjunto de métodos
// através das interfaces podemos utilizar polimorfismo.

// O que a interface Geometry faz? Implementa os métodos para medir área e perimetro.

type geometry interface {
	area() float64
	perim() float64
}

// Dados que serão utilizados pelo método são declarados dentro da struct
type rect struct {
	width, height float64
}

// Dados que serão utilizados pelo método são declarados dentro da struct
type circle struct {
	radius float64
}

// Métodos da interface. declaro que 'r' é do tipo rect e o método que desejo utilizar mais o retorno deste método.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// função que recebe os valores de rect e circle para implementar sueus métodos
// func meansure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

func main() {

	n := []geometry{rect{5, 6}, circle{5}}

	for _, i := range n {

		fmt.Println("\n", i.area(), "\n", i.perim())
	}

}
