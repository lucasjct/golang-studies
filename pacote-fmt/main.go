package main

// documentação:
// https://pkg.go.dev/fmt

import "fmt"

func main() {

	x := "Olá"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y) // função de sprint

	//fmt.Print(x)
	// fmt.Println(x) // Print line add um linha nova no final
	// fmt.Print(y)

	fmt.Println(z)

}
