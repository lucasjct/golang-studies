package main

import "fmt"

func main() {
	niveltocansado := 8

	if niveltocansado == 1 {
		fmt.Println("Queria deitar na cama")
	} else if niveltocansado == 2 {
		fmt.Println("Levanta")
	} else {
		fmt.Print("Estuda!")
	}
}
