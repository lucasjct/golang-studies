package main

import "fmt"

type exercicio int

var x exercicio

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var x = 42
	fmt.Printf("%v", x)
}
