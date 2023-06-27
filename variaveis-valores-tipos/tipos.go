package main

import "fmt"

// variável package level scope, Precisa do var e ser fora do code block
// tipos primitivos em Go: integer, string e boolean
// Tipos de dados compostos: são compostos de tipos primitivos, e criados pelo usuário. São:
// slice, array, struct, map
var y = 10

var a int
var b bool
var c float32
var d string

// go é uma linguagem de tipagem estática. Ou seja, seu tipo declarado não pode receber valor de outro tipo que o original
func main() {
	// O operador curto := serve para declarar variável (Operador de declaração) e atribuí-la.
	// O Operador curto se assemalha ao Gopher, mascote do Go.

	// Os operadores curtos podem ser utilizados apenas dentro de __code blocks__.

	// Tipagem Automática:

	// O operador curto decifra o tipo da variável de acordo
	// com o valor que foi informado.

	// Expressão é tudo que retorna um resultado.A combinação de um ou mais valores, constantes, variáveis, operadores e funções que a linguagem interpreta e usa para produzir outro valor.

	// Statement (Declaração) é uma linha de código. É o menor elemento de um programa que expressa uma ação a ser executada;
	//

	x := 10
	y := "olá boa noite"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

	// Operador de Atribuição, diferente da marmota que é o operador de declaração
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)

	retornarLevelScope()
}

// função para demonstrar que posso acessar uma variável do package level scope
func retornarLevelScope() {
	fmt.Println(y)
}

// Keywords Go Lang: https://go.dev/ref/spec#Keywords

// Diferença entre Declaração, Inicialização e Atribuição
// Declaração => separa um endereço de memória para a variável
// Inicialização => primeiro valor que a variável declarada recebe
// Atribuição => Definição de valor para um variável

// Valor 0
// em Go a variável é inicializada com valor 0 quando não atribuímos valores. Cada tipo na linguagem recebe um valor 0 diferente.
// - int: 0
// - floats: 0.0
// - booleans: false
// - strings: ""

// Usar sempre que possível o := para atribuição
