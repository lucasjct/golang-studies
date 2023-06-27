package main

import "fmt"

var x interface{}

func main() {

	// x := 5

	// switch true {

	// case x < 4:
	// 	fmt.Println("chis é menor que 4")
	// case x == 3:
	// 	fmt.Println("chis é igual a 3")
	// case x == 5:
	// 	fmt.Println("chis é igual a 5")
	// case x > 1:
	// 	fmt.Println("chis é maior que 1")

	// }

	//! Exemplo Dois:

	// quemTaNoEscritorioHoje := "Marina"

	// switch quemTaNoEscritorioHoje {
	// case "Zezinho":
	// 	fmt.Println("Quem está no escritório Zezinho.")
	// pula a linha acima e vai para a debaixo
	// 	fallthrough
	// case "Marquinhos":
	// 	fmt.Println("Quem está no escritório Marquinhos.")
	// case "Ricardinho":
	// 	fmt.Println("Quem está no escritório Ricardinho.")
	// com default, é ofericida uma resposta se nenhuma das condições forem atendidas.
	// default:
	// 	fmt.Println("Escritório vazio")
	// }

	//! Exemplo três. Case composto

	// numeros := 4

	// switch numeros {
	// case 1, 2:
	// 	fmt.Println("Número um e dois")
	// case 3, 4:
	// 	fmt.Println("Numeros três e quatro")

	// }

	//! Exemplo com interface vazia

	x = 1
	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case float64:
		fmt.Println("é um float")
	case bool:
		fmt.Println("é um boolean")
	case string:
		fmt.Println("é uma string")
	}

}
