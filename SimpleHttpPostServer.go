package main

import (
	"bytes"
	"net/http"
)

func main() {
	http.HandleFunc("/", printBody)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		println("网络错误")
	}
}

func printBody(write http.ResponseWriter, request *http.Request) {
	write.WriteHeader(200)
	buf := new(bytes.Buffer)
	buf.ReadFrom(request.Body)
	s := buf.String()
	println(s)
}
