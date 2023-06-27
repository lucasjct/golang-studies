### Comandos Go   
***


* `go version`   
Retorna a versão do Go    

* `go env`  
Retorna as variáveis de ambiente do go. É importante verificar essas varíveis para ver se o go path está correto.   

* `go help + <parametro do comando>`   
Retorna as informações do comando.   

* `go run`
    * `go run <file name>`  
    
    *  `go run *.go`  
    Executa todos os arquivos . 

* `go build`  
gera um arquivo binário.   

* `go install`   
Instala o programa no path. Então quando digita o nome do programa no terminal, esse programa é executado por estar instalado no $GOPATH/bin.   

* ` -race`   
Go aceita flags.    


### Compilação cruzada    

Compilar programa Go para ser executado em diferentes Sistemas Operacionais. Para isso, temos que saber duas variáveis:    

* GOOS (sistema operacional em que o programa será executado)   
* GOARCH (arquitura do processador em que o programa será executado).    


Dado o pequeno programa abaixo:   

``` go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
```   
Para fazer a compilação cruzada e executar o programa compilado no Windows, precisamos digitar o seguinte comando no terminal.   

`GOOS=windows GOARCH=amd64 go build main.go`

* Para o Linux:  
`GOOS=linux GOARCH=amd64 go build main.go`   

* Para o Windows   
`GOOS=darwin GOARCH=amd64 go build main.go`



