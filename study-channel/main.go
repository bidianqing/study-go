package main

import (
	"fmt"
	"sync"
	"time"
)

// channel本质是队列，先进先出，线程安全
// 引用类型
// 在没有使用协程的情况下，写入数据和消费数据会改变管道的长度，管道的长度不能超过容量
// 在没有使用协程的情况下，如果管道的长度是0，就不能从管道里取数据，否认报错
// 使用内置函数close可以关闭管道，当管道关闭后，就不能再向管道写数据了，但是仍然可以从该管道读取数据。
// 在没有使用协程的情况下，遍历管道前要进行管道的关闭
// 发生阻塞的情况
// 1、无缓冲区时单独的读或者单独的写
// 2、缓冲区为空时进行读数据
// 3、缓冲区满时进行写数据

func main() {
	//TestChan()

	//TestCloseChan()

	// TestEachChan()

	//TestGoRoutineAndChan()

	//TestSelectChan()
}

func TestSelectChan() {
	c1 := make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 10
	}()

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "hello golang"
	}()

	select {
	case v := <-c1:
		fmt.Println("从c1获取到:", v)
	case v := <-c2:
		fmt.Println("从c2获取到:", v)
	}
}

func TestGoRoutineAndChan() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("写入数据成功", i)
			time.Sleep(time.Second)
		}
		close(c)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range c {
			fmt.Println("读取数据成功", v)
		}
	}()

	wg.Wait()

	fmt.Println("执行成功")
}

func TestEachChan() {
	c := make(chan int, 3)
	c <- 10
	c <- 11
	c <- 12

	close(c) // 未使用协程时，遍历管道前要进行管道的关闭
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("管道长度为", len(c), "，容量为", cap(c))
}

func TestCloseChan() {
	c := make(chan int, 3)

	c <- 10
	close(c)
	c <- 11 // 管道关闭以后就不能写入数据了，此行报错
	<-c
}

func TestChan() {
	c := make(chan int, 3) // 容量为3
	fmt.Println("管道长度为", len(c), "，容量为", cap(c))

	c <- 10
	c <- 11
	c <- 12

	<-c // 如果不消费的话，是不能继续写入的
	c <- 13

	<-c
	<-c
	<-c

	fmt.Println("管道长度为", len(c), "，容量为", cap(c))

	//<-c

	fmt.Println(c)
}
