package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// os tipos abaixo são slices de carros
type ordernarPorPotencia []carro

// com a funções abaixo, posso customizar os sorts, utilizando o método sort.Sort(ordernarPorPotencia(carros))
func (x ordernarPorPotencia) Len() int           { return len(x) }
func (x ordernarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (x ordernarPorPotencia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordernarPorConsumo []carro

func (x ordernarPorConsumo) Len() int           { return len(x) }
func (x ordernarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (x ordernarPorConsumo) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordernarPorLucroDonoDoPosto []carro

func (x ordernarPorLucroDonoDoPosto) Len() int           { return len(x) }
func (x ordernarPorLucroDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (x ordernarPorLucroDonoDoPosto) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	carros := []carro{
		{"Chevete", 50, 8},
		{"Porshe", 300, 5},
		{"Fusca", 20, 30},
	}

	fmt.Println(carros)
	sort.Sort(ordernarPorPotencia(carros))
	fmt.Println("", carros)

	sort.Sort(ordernarPorConsumo(carros))
	fmt.Println("", carros)

	sort.Sort(ordernarPorLucroDonoDoPosto(carros))
	fmt.Println("", carros)

}
