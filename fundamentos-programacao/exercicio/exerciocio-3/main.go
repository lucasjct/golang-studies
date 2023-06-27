package main

import "fmt"

func main() {

	const num float64 = 10 // tipada
	const num2 = 10        // não tipada

	fmt.Printf("%v\n%T ", num, num)
	fmt.Printf("%v\n%T ", num2, num2)

}
