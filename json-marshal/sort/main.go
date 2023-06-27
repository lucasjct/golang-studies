package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"estandate", "e", "desfraudado", "preto", "e", "branco", "é", "a", "sua", "cor"}
	is := []int{8, 001, 00, 0, 1, 89, 336, 6, 11, 258}

	stringSort(ss)
	intSort(is)

}

func stringSort(ss []string) {

	sort.Strings(ss)
	fmt.Println(ss)

}

func intSort(is []int) {

	sort.Ints(is)
	fmt.Println(is)

}
