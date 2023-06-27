package main

import (
	"encoding/json"
	"fmt"
)

//https://transform.tools/json-to-go

// Documentação
// https://pkg.go.dev/encoding/json#Unmarshal

type Pessoa struct {
	Nome      string  `json:"nome"` // esses campos são chamados de tags: https://pkg.go.dev/encoding/json#Marshal
	Sobrenome string  `json:"sobrenome"`
	Idade     int     `json:"idade"`
	Profissao string  `json:"profissao"`
	Salario   float64 `json:"salario"`
}

func main() {

	sb := []byte(`{"Nome":"Ana Maria","Sobrenome":"Carvalho Teixeira","Idade":73,"Profissao":"Aposentada","Salario":1265.5}`)
	var p Pessoa
	err := json.Unmarshal(sb, &p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(p)

}
