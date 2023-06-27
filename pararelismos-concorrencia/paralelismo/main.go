// A race condition
// occurs when multiple threads try to access and modify the same data (memory address)

// o resultado deste código é não ter o total exato de goroutines por conta do problema de race condition.
// como diferentes processos tentarão acessar e modificar o mesmo dado na memória, a incremetação não será correta.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	totalDeGoRoutines := 10
	contador := 0

	// experar a execução da função
	var wg sync.WaitGroup

	wg.Add(totalDeGoRoutines)

	for x := 0; x < totalDeGoRoutines; x++ {

		go func() {

			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()

	}
	// esperar as go functions anonimas terminar de serem executadas dentro do for.
	wg.Wait()

	fmt.Println(contador)

}
