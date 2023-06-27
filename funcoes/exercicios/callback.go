package main

import "fmt"

func processInput(input int, callback func(int)) {

	result := input * 2
	callback(result)
}

func myCallback(result int) {

	fmt.Println("resultado", result)
}

func main() {

}
