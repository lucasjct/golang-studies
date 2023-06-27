package main

import "fmt"

func main() {

	num := 10
	fmt.Printf("%d\t%b\t%#X\n", num, num, num)

	numDes := num << 1
	fmt.Printf("%b\t %d", numDes, numDes)
}
