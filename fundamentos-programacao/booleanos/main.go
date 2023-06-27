package main

import "fmt"

var x bool

func main() {
	// Boleanos são true, false, 1 e 0
	// São fundamentais para estruturas condicionais com if ou switch

	fmt.Println("zero value:", x) // zero value
	x := true

	// os parentesses não são necessários. Servem apenas para melhorar a leitura.
	fmt.Println(x)
	x = (10 > 100)
	y := (10 == 100)
	z := (10 < 100)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
