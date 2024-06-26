package main

import (
	"fmt"
	"sync"
	"time"
)

// channel本质是队列，先进先出，线程安全
// 引用类型
// 管道关闭以后不能重复关闭，否认会panic: close of closed channel
// 管道关闭以后就不能往里面写数据，否则会panic: send on closed channel
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

	//TestEachChan()

	//TestGoRoutineAndChan()

	TestSelectChan()
}

// select 是 Go 语言中的一种控制结构，用于在多个通信操作中选择一个可执行的操作。它可以协调多个 channel 的读写操作，使得我们能够在多个 channel 中进行非阻塞的数据传输、同步和控制。
// Go 语言中的 select 语句是一种用于多路复用通道的机制，它允许在多个通道上等待并处理消息。相比于简单地使用 for 循环遍历通道，使用 select 语句能够更加高效地管理多个通道。
func TestSelectChan() {
	c1 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c1 <- i
			time.Sleep(time.Second * 5)
		}
		close(c1)
	}()

	c2 := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c2 <- "hello golang"
			time.Sleep(time.Second * 2)
		}
		close(c2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case v, ok := <-c1:
				if ok {
					fmt.Println("从c1获取到:", v)
				} else {
					c1 = nil
				}
			case v, ok := <-c2:
				if ok {
					fmt.Println("从c2获取到:", v)
				} else {
					c1 = nil
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("执行成功")
}

func TestGoRoutineAndChan() {
	c := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("写入数据成功", i)
			fmt.Println("管道长度为", len(c), "，容量为", cap(c))
			time.Sleep(time.Second)
		}
		close(c)
	}()

	wg.Add(1)
	go func() {

		// for {
		// 	v, ok := <-c
		// 	fmt.Println("读取数据成功", v, ok)
		// 	if !ok {
		// 		fmt.Println("ok=", ok)
		// 		wg.Done()
		// 		break
		// 	}
		// }

		defer func() {
			fmt.Println("执行defer")
			wg.Done()
		}()
		// 在使用协程的情况下，如果管道关闭，在消费完管道里的数据以后，自动结束循环
		for v := range c {
			fmt.Println("读取数据成功", v)
			time.Sleep(time.Second * 5)
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
	c := make(chan int, 2)
	c <- 10
	c <- 11
	close(c)
	//close(c) // 管道关闭以后不能重复关闭，否认会panic: close of closed channel
	// c <- 12 // 管道关闭以后就不能往里面写数据，否则会panic: send on closed channel

	num, ok := <-c
	fmt.Println(num, ok)

	num2, ok2 := <-c
	fmt.Println(num2, ok2)

	num3, ok3 := <-c
	fmt.Println(num3, ok3)
}

func TestChan() {
	c := make(chan int, 3) // 容量为3
	fmt.Println("管道长度为", len(c), "，容量为", cap(c))

	c <- 10
	fmt.Println("管道长度为", len(c), "，容量为", cap(c))
	c <- 11
	c <- 12

	<-c // 如果不消费的话，是不能继续写入的
	c <- 13

	<-c
	<-c
	<-c

	fmt.Println("管道长度为", len(c), "，容量为", cap(c))

	//<-c
}
