package main

import "fmt"

func main() {

	ss := [][]int{

		//   0  1  2
		{1, 2, 3}, // Índice 0
		{4, 5, 6}, // Índice 1
		{7, 8, 9}, // Índice 2
	}

	fmt.Println(ss[1][2])
	fmt.Println(ss)

	// devemos enxergar o slice multi-dimensional como se fosse uma planilha do excel.
	// é um slice de slices.

}
