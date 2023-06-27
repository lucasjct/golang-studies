package main

import "fmt"

func main() {

	esporteFavorito := "futebol"

	switch {
	case esporteFavorito == "futebol":
		fmt.Println("Pooooonnteeeeee")
	case esporteFavorito == "Tênis":
		fmt.Println("Só conheço o Guga")
	case esporteFavorito == "Basquete":
		fmt.Println("Oscar")

	default:
		fmt.Println("Preciso começar a praticar esporte")
	}
}
