package main

import (
	"fmt"
	"github.com/awesomerevert/JSParser"
	"net/http"
	"time"
)

func init() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		jsFilePath := `public/test.js`
		rt, _ := JSParser.JsParser(jsFilePath, "a", 1, 2)
		fmt.Fprintln(writer, rt)
	})

	http.HandleFunc("/bs", func(writer http.ResponseWriter, request *http.Request) {
		jsFilePath := `public/example.js`
		timeStamp := time.Now().UnixNano()
		result := fmt.Sprintf("%v", timeStamp)
		fmt.Println(result)
		if rt, err := JSParser.JsParser(jsFilePath, "encodeInp", result); err != nil {
			fmt.Printf(`Error Js Parse call encodeInp`)
		} else {
			fmt.Fprintln(writer, rt)
		}
	})
}

func main() {
	if err := http.ListenAndServe(`:9090`, nil); err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
