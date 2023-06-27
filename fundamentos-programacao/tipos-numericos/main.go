package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "香"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a) // 1 byte
	e := []byte(b) // 2 bytes
	f := []byte(c) // 3 bytes

	fmt.Printf("%v, %v, %v\n", d, e, f)

	VerificarProcessadorESo()
	Overflow()

}

// Vertificar sistema operacioal e arquitetura do processador
func VerificarProcessadorESo() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

// Função para demostrar quando ultrapassmos o limite de caracteres de um tipo uint69. Isso é chamado de Overflow

func Overflow() {

	var i uint16
	i = 65535
	fmt.Println(i)
	i++ // na variável contadora, quando ultrapassa o valor limite, ele retonar para zero e recomeça a contagem.
	fmt.Println(i)
	i++
	fmt.Println(i)
}
