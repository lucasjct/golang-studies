package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	h := crc32.NewIEEE()

	h.Write([]byte("Lucas José Carvalho Teixe"))

	v := h.Sum32()
	fmt.Println(v)

}
