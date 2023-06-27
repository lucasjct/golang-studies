package main

import "fmt"

//! maps é uma estrutura de dados organizada entre chaves e valores.

func main() {

	amigos := map[string]int{

		"Alfredo": 5512255,
		"Joana":   66785455,
	}

	//! add dados no map

	amigos["gopher"] = 4444455

	fmt.Println(amigos)
	fmt.Println(amigos["Joana"])

	//! comma ok idiom

	if sera, ok := amigos["fantasma"]; ok {
		fmt.Println(sera)
	} else {
		fmt.Println("Não tem!")
	}

	demoMap()

}

//! Como se deleta e utiliza o range em maps

func demoMap() {

	convidados := map[int]string{

		1: "Ricardo",
		2: "Marina",
		3: "Ana",
	}

	fmt.Println(convidados)

	for key, value := range convidados {
		fmt.Println(key, value)
	}

	//! Para deletar
	delete(convidados, 1)
	fmt.Println(convidados)
}
