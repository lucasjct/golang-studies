// Strings - methods from book Introducing Go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// constains (s, substr string) bool

	name := "batata"

	contains(name)
	contarStrings(name)
	verificarPrefixo(name)
	temSufixo(name)
	localizarPosicao(name)
	unirStrings()
	repetir("a", 8)
	replaceString()
	stringToList(name)
	minusculas(name)
	maiuscula(name)
}

// Procurar uma string menor dentro de uma string maior
func contains(name string) {
	fmt.Println(strings.Contains(name, "a"))
}

// contar o numero de vezes que uma string ocorre dentro de outra maior
func contarStrings(name string) {
	fmt.Println(strings.Count(name, "s"))
}

// Verificar se uma string maior começa com determinado caractere
func verificarPrefixo(name string) {
	fmt.Println(strings.HasPrefix(name, "a"))
}

// Verificar se uma string maior termina com determinado caractere
func temSufixo(name string) {
	fmt.Println(strings.HasSuffix(name, "o"))
}

// localizar a posição de uma string. a contagem inicia do 0
func localizarPosicao(name string) {
	fmt.Println(strings.Index(name, "a"))
}

// unir uma lista de strings em uma única string
func unirStrings() {
	fmt.Println(strings.Join([]string{"c", "a", "s", "a"}, ""))
}

// repetir uma srting
func repetir(s string, q int) {
	fmt.Println("string was repeted ", q, " times: ", strings.Repeat(s, q))

}

// fazer replace dentro de uma strings
func replaceString() {
	fmt.Println(strings.Replace("aaaa", "a", "ti", 2))
}

// transformar string em lista
func stringToList(name string) {

	fmt.Println(strings.Split(name, ""))
}

// converter para minúsculas
func minusculas(name string) {
	fmt.Println(strings.ToLower(name))
}

// converter para maiúsculas
func maiuscula(name string) {
	fmt.Println(strings.ToUpper(name))
}
