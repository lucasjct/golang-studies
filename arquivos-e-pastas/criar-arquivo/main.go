package main

import "os"

func main() {

	file, err := os.Create("text.html")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("<h1>Hello Word</h1>")
}
