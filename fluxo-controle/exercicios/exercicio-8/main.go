package main

import (
	"fmt"
)

func main() {

	nivelAlegria := 11

	switch {
	case nivelAlegria > 5 && nivelAlegria <= 10:
		fmt.Println("Uhull")

	case nivelAlegria <= 5:
		fmt.Println("Vai melhorar")

	default:
		fmt.Println("sem condiçoes")
	}

}
