package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 100
	testReflect(num)
}

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println("reType:", reType)

	reValue := reflect.ValueOf(i)
	fmt.Println("reValue:", reValue)

	fmt.Println(1 + reValue.Int())

	i2 := reValue.Interface()
	n := i2.(int)
	fmt.Println(n)
}
