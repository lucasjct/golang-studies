package main

import "fmt"

type exercicio int

var x exercicio
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Printf("%v\n", x)

	fmt.Println("==============")

	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}
