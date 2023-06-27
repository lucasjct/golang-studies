package main

import (
	"fmt"
	"math/rand"
)

func main() {

	ss := []string{"Calabresa", "Marguerita", "Mussarela", "Frango com catupiry"}
	randomChoice(ss)

}

func randomChoice(ss []string) {

	fmt.Println(ss[rand.Intn(len(ss))])
}
