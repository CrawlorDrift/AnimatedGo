//package main
//
//import (
//	"fmt"
//	"github.com/robertkrimen/otto"
//	"io/ioutil"
//)
//
//func JsParser(filePath string, functionName string, args ...interface{}) (result string) {
//	bytes, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		panic(err)
//	}
//	vm := otto.New()
//	_, err = vm.Run(string(bytes))
//	if err != nil {
//		panic(err)
//	}
//	value, err := vm.Call(functionName, nil, args...)
//	if err != nil {
//		panic(err)
//	}
//	return value.String()
//}
//
//func main() {
//	filePath := "public/example.js"
//	//先读入文件内容
//	bytes, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		panic(err)
//	}
//	vm := otto.New()
//	_, err = vm.Run(string(bytes))
//	if err != nil {
//		fmt.Printf("Error, RUN : %v", err)
//	}
//
//	data := "asd"
//	//encodeInp是JS函数的函数名
//	value, err := vm.Call("encodeInp", nil, data)
//	if err != nil {
//		fmt.Printf("Error, Call : %v", err)
//	}
//	fmt.Println(value.String())
//
//}

package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
	"time"
)

// JsParser is a load js parse function, you can use it to call js get result
func JsParser(filePath string, functionName string, args ...interface{}) (result string, err error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error, when readfile: %v\n", err)
		return "", err
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err != nil {
		fmt.Printf("error, when parse js file :%v\n", err)
		return "", err
	}
	value, err := vm.Call(functionName, nil, args...)
	if err != nil {
		fmt.Printf("error, when call js function :%v\n", err)
		return "", err
	}
	return value.String(), err
}

func init() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		jsFilePath := `public/test.js`
		rt, _ := JsParser(jsFilePath, "a", 1, 2)
		fmt.Fprintln(writer, rt)
	})

	http.HandleFunc("/bs", func(writer http.ResponseWriter, request *http.Request) {
		jsFilePath := `public/example.js`
		timeStamp := time.Now().UnixNano()
		result := fmt.Sprintf("%v", timeStamp)
		fmt.Println(result)
		if rt, err := JsParser(jsFilePath, "encodeInp", result); err != nil {
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
