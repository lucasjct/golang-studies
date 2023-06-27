package main

import "fmt"

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês", mes)

		for dias := 1; dias <= 30; dias++ {
			fmt.Print(" ", dias)
		}
		fmt.Println("\n")
	}

}
