// A race condition
// occurs when multiple threads try to access and modify the same data (memory address)

// o resultado deste código é não ter o total exato de goroutines por conta do problema de race condition.
// como diferentes processos tentarão acessar e modificar o mesmo dado na memória, a incremetação não será correta.

// O exemplo abaixo implementa o mutex para contornar o problema de race condition.

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

	// variavel para instanciar os métodos do mutex.
	var mu sync.Mutex

	for x := 0; x < totalDeGoRoutines; x++ {

		go func() {

			// tranca um trecho de código e faz com que somente um trhead (somente uma goroutine), possa executar por vez.
			// As demais treads terão que esperar o fim de uma execução para poderem executar
			mu.Lock()

			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()

			mu.Unlock()
		}()

		fmt.Println(runtime.NumGoroutine())

	}
	// esperar as go functions anonimas terminar de serem executadas dentro do for.
	wg.Wait()
	fmt.Println(contador)
	fmt.Println(runtime.NumGoroutine())

}
