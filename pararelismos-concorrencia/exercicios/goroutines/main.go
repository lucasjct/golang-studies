package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	novaGoRoutine(2)
	wg.Wait()

}

func novaGoRoutine(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {

		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine numero", i+1)
			wg.Done()
		}(x)
	}
}
