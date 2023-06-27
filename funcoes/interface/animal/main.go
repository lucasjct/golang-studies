package main

import "fmt"

type animal interface {
	speak() string
}

type dog struct {
}

func (d dog) speak() string {
	return "au-au-au"
}

type cat struct {
}

func (c cat) speak() string {
	return "miau miau"
}

type human struct {
}

func (h human) speak() string {
	return "Howww man!"
}

func main() {
	a := []animal{dog{}, cat{}, human{}}

	for _, i := range a {
		fmt.Println(i.speak())
	}

}
