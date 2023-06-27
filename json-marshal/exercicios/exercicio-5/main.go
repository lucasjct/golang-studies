package main

//- Partindo do código abaixo, ordene os []user por idade e sobrenome.

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user

// Métodos da ordenação por idade
func (x ordenarPorIdade) Len() int           { return len(x) }
func (x ordenarPorIdade) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x ordenarPorIdade) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorSobrenome []user

// Métodos da ordenação por sobrenome
func (x ordenarPorSobrenome) Len() int           { return len(x) }
func (x ordenarPorSobrenome) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (x ordenarPorSobrenome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   85,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// your code goes here
	fmt.Println("Ordenada por sobrenome")
	sort.Sort(ordenarPorSobrenome(users))
	Ordernacao(users)
	//fmt.Println(users)

	fmt.Println("Ordenada por idade")
	sort.Sort(ordenarPorIdade(users))
	Ordernacao(users)

}

func Ordernacao(user []user) {

	for k, v := range user {

		fmt.Println(k, "\tAge:", v.Age, "\tNome completo:", v.First, v.Last)
		fmt.Print("\n")

		for _, va := range v.Sayings {
			fmt.Println(va)
		}
		fmt.Println("")
	}

	fmt.Println("")

}
