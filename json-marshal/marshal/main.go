package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	Salario   float64
}

func main() {

	paulo := Pessoa{"Paulo", "Teixeira", 73, "Aposentado", 3685.74}
	ana := Pessoa{"Ana Maria", "Carvalho Teixeira", 73, "Aposentada", 1265.50}

	// como o método marshal, podemos ordenar as informações da struct e trasformar e json.
	// marshal transforma bytes em json, ele torna o objeto serializável para poder ser utilizado/transmitido
	p, err := json.Marshal(paulo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p))

	p2, err := json.Marshal(ana)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p2))

}
