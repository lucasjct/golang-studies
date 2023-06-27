// IOTA.
package main

import "fmt"

func main() {

	const (
		_ = iota + 2023
		x
		y
		z
		a
	)

	fmt.Println(x, y, z, a)
}
