package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	http.Handle(
		"assets/index.html",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)

}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		<head><title>Hello Word</title></head>
		<body>
		Hello, Word!
		</body>
		</html>`,
	)
}
