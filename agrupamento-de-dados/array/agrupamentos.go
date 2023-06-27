package main

import "fmt"

func main() {

	// - Arrays
	// - Slices // Bloco que foi construído sobre o Array
	// - Maps
	// - Structs

	var x [5]int

	x[0] = 1
	x[1] = 3
	x[3] = 4

	fmt.Println(x[0], x[1], x[3])
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))

}
