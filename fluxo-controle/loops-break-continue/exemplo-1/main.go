package main

// O continue dará sequência ao loop após averiguar determinada condição e somente respeitará a condição.
// O break destruíra o loop após averiguar determina condição.

func main() {

	for i := 0; i <= 20; i++ {
		if i == 3 || i == 4 {
			continue
		}
		print(" ", i)
	}

}
