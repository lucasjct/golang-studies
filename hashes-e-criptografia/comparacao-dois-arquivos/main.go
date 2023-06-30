package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h1, err := getHash("text1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("text2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

}

func getHash(file string) (uint32, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)

	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
