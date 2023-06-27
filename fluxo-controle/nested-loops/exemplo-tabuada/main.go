package main

import "fmt"

func main() {

	for tab := 0; tab < 11; tab++ {

		fmt.Println("Tabuado do: ", tab)
		for x := 0; x < 11; x++ {
			fmt.Print(" ", x*tab)
		}

		fmt.Print("\n")

	}
}
