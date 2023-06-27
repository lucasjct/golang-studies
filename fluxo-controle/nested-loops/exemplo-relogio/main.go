// repetições hierárquicas
package main

import "fmt"

func main() {
	for hours := 0; hours <= 12; hours++ {
		fmt.Print("Hora ", hours, ": ")
		for i := 0; i < 60; i++ {
			fmt.Print(" ", i)

		}
		fmt.Println("\n")

	}
}
