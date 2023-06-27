package main

import "fmt"

func main() {

	hobbies := map[string][]string{

		"teixeira_paulo": {"Jogar bola", "escutar música"},
		"ana_teixeira":   {"ficar com o Gael", "ficar com os filhos"},
		"arlindo_cruz":   {"fazer samba", "ver os amigos e a família"},
	}

	delete(hobbies, "arlindo_cruz")

	for k, v := range hobbies {
		fmt.Println(k)
		for _, i := range v {
			fmt.Println("\t", i)

		}
		fmt.Println("")

	}
}
