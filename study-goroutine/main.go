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
// go语言在执行goroutine的时候是没有返回值的，这时候我们要用到go语言中特色的channel来获取返回值

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Hello Java" + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Hello Golang" + strconv.Itoa(i))
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()

	fmt.Println("执行完成")
}
