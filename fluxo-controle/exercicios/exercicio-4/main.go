package main

import "fmt"

func main() {

	anoquenasci := 1989
	anoatual := 2023

	// for {
	// 	if anoquenasci <= anoatual {

	// 		fmt.Println(anoquenasci)
	// 		anoquenasci++

	// 	} else {
	// 		break
	// 	}

	// }

	//! solução diferente

	for {
		if anoquenasci == anoatual {
			break
		}
		fmt.Println(anoquenasci)
		anoquenasci++

	}

}
