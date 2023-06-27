package main

// Diferente do marshal que guarda o resultado da função na variável
// o Encoder exibe o resultado dá saída do terminal
// O encoder é uma interface abstrata que permite a codificação de dados

import (
	"encoding/json"
	"os"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	Salario   float64
}

func main() {

	p := Pessoa{"Lucas", "Teixeira", 33, "TI", 80000.00}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(p)

}
