// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

package main

import "fmt"

type area interface {
	AreaCirculo() float64
	AreaQuadrado() int
}

type Quadrado struct {
	ladoA, ladoB int
}

type Circulo struct {
	raio float64
}

type figura struct {
}

func (q Quadrado) AreaQuadrado() int {
	return q.ladoA * q.ladoB
}

func (c Circulo) AreaCirculo() float64 {
	return 2 * 3.14 * c.raio

}

func main() {
	q := Quadrado{10, 10}
	c := Circulo{10}
	fmt.Println(q.AreaQuadrado())
	fmt.Println(c.AreaCirculo())
}
