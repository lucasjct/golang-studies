package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	base64file "github.com/lucasjct/curso-go/cold-start/convert-to-string"
)

func SendRequest() {
	url := "https://yfpc3a5779.execute-api.us-east-1.amazonaws.com/hml/api/tab"
	method := "POST"

	filebase64 := base64file.ConvertToString()
	payload := strings.NewReader(filebase64)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Para o bearer token, podemos criar uma função de login o retorna
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6Imx1Y2FzamN0ZWl4ZWlyYUBnbWFpbC5jb20iLCJqdGkiOiIzMDlkYzdjOC1mNmQ0LTQ0ZDctOWZmOC0yOTI1MTZkOTc4YTYiLCJJZFVzdWFyaW8iOiIxNSIsIk5vbWUiOiJMVUNBUyBURVNURVMiLCJFbWFpbCI6Imx1Y2FzamN0ZWl4ZWlyYUBnbWFpbC5jb20iLCJUcm9jYXJTZW5oYSI6IkZhbHNlIiwiQW1iaWVudGUiOiJWZWljdWxvc2JhY2tvZmZpY2UiLCJuYmYiOjE2ODAxMDQ4MzAsImV4cCI6MTY4MDExNTYzMCwiaWF0IjoxNjgwMTA0ODMwfQ.fF-_PFSHLgRiWOiMqqEclvi6Z_h5a9IBq2QASYw73sA")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
