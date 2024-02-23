package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// 使用go关键字开启一个协程
	// 主死从随
	// 使用WaitGroup控制协程退出
	// go语言在执行goroutine的时候是没有返回值的，这时候我们要用到go语言中特色的channel来获取返回值
	//TestGoRoutine()

	// 互斥锁机制，防止多个协程同时操作一个资源
	// 互斥锁性能和效率比较低
	//TestMutex()

	// 读写锁 经常用于读多写少的场景
	// 打印修改数据后必须紧跟修改数据成功，期间不可能出现读取操作
	//TestRWMutex()
}

func TestRWMutex() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(&wg, &lock)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(&wg, &lock)
	}

	wg.Wait()

	fmt.Println("执行成功")
}

func read(wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.RLock()
	fmt.Println("读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write(wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	lock.Lock()
	fmt.Println("修改数据")
	time.Sleep(time.Second * 2)
	fmt.Println("修改数据成功")
	lock.Unlock()
}

func TestMutex() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var num int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			lock.Lock()
			num++
			lock.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			lock.Lock()
			num--
			lock.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("num的值为", num)
}

func TestGoRoutine() {
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
