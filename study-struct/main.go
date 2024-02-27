package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Id   int64
	Name string
	B    bool
}

func main() {
	fmt.Printf("bool:   %d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("string: %d\n", unsafe.Alignof(string("a")))
	fmt.Printf("int:    %d\n", unsafe.Alignof(int(0)))
	fmt.Printf("int32:  %d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("int64:  %d\n", unsafe.Alignof(int64(0)))
	fmt.Printf("float64:  %d\n", unsafe.Alignof(float64(0)))
	fmt.Printf("float32:%d\n", unsafe.Alignof(float32(0)))

	fmt.Println(unsafe.Sizeof(Person{}))
}
