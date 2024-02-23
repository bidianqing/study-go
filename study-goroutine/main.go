package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 使用go关键字开启一个协程
// 主死从随
// 使用WaitGroup控制协程退出

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)
	}

	wg.Add(1)
	go PrintJava()

	wg.Add(1)
	go PrintGolang()

	wg.Wait()

	fmt.Println("执行完成")
}

func PrintJava() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Java" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

func PrintGolang() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Golang" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
