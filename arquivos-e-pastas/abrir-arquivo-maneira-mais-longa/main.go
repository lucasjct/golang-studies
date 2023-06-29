package main

import (
	"fmt"
	"os"
)

// maneira menos concisa de abrir um arquivo

func main() {
	file, err := os.Open("tests.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return

	} else {

		fmt.Println(stat)
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
