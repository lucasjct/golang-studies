package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// método de criptografia do tipo hash para senhas

func main() {

	senha := "20julho1980"
	//senhaerrada := "20julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	err = bcrypt.CompareHashAndPassword(sb, []byte(senha))
	if err != nil {
		fmt.Println(err)
	}

}
