package JSParser

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
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