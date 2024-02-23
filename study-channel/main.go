package main

import "fmt"

func main() {
	// channel本质是队列，先进先出，线程安全
	// 引用类型
	// 写入数据和消费数据会改变管道的长度，管道的长度不能超过容量
	// 在没有使用协程的情况下，如果管道的长度是0，就不能从管道里取数据，否认报错
	// 使用内置函数close可以关闭管道，当管道关闭后，就不能再向管道写数据了，但是仍然可以从该管道读取数据。
	// 遍历管道前要进行管道的关闭

	//TestChan()

	//TestCloseChan()

	//TestEachChan()
}

func TestEachChan() {
	c := make(chan int, 3)
	c <- 10
	c <- 11
	c <- 12

	close(c) // 遍历管道前要进行管道的关闭
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
	c := make(chan int, 3)
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
