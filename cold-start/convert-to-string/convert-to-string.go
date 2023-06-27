package base64file

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ConvertToString() string {
	jsonFile, err := os.Open("payload/base64.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// tratar base64 como arquivo de bytes e converter para string
	var buf bytes.Buffer
	io.Copy(&buf, jsonFile)
	payload := string(buf.String())

	return payload
}
