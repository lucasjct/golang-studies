package main

import "fmt"

func main() {
	nome := `package main

	import "fmt"
	
	func main() {
		nome := "fmt.Println()"
	
		fmt.Println(nome)
	}
	`

	fmt.Println(nome)
}
