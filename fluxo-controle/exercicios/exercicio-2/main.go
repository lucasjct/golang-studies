package main

import "fmt"

func main() {

	for i := 65; i <= 90; i++ {

		fmt.Println(i)

		for x := 1; x <= 3; x++ {
			c := i
			asciiValue := rune(c)
			fmt.Printf("%#U %c\n", c, asciiValue)
		}
		fmt.Println("")
	}

}
