package main

import (
	"bytes"
	"flag"
	"net/http"
)

func main() {
	// 从命令行获取参数
	var host string
	var port string

	flag.StringVar(&host, "h", "0.0.0.0", "host")
	flag.StringVar(&port, "p", "8080", "端口")

	flag.Parse()

	http.HandleFunc("/", printBody)
	err := http.ListenAndServe(host+":"+port, nil)
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
