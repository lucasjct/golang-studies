package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {

	// quando temos o printf, precisamos utilizar os caracteres literal para formatar a saída.
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)

}
