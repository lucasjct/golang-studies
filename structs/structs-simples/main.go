package main

import "fmt"

// armazenar dados com tipos diferentes

func main() {

	type cliente struct {
		nome      string
		sobrenome string
		fumante   bool
	}

	clienteA := cliente{
		nome:      "Lucas",
		sobrenome: "Teixeira",
		fumante:   true,
	}

	clienteB := cliente{"Cássio", "Gabriel", false}

	fmt.Println(clienteA)
	fmt.Println(clienteB)

}
