package main

import (
	"fmt"
	"reflect"
)

func test(foo interface{}) {
	toCall := reflect.ValueOf(foo)
	toCall.Call([]reflect.Value{})
}

func test2() {
	fmt.Println("hey")
}
func main() {
	test(test2)

}
