package main

import "fmt"

func main() {

	// strings são imutáveis

	a := "e"
	b := "é"
	c := "香"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v", d, e, f)
}

// artigo sobre strings: https://go.dev/blog/strings
