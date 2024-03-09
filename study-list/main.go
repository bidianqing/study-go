package main

import (
	"container/list"
)

// 双向链表

func main() {
	l := list.New()

	l.PushBack("A")
	l.PushBack("B")
	l.PushBack("C")
}
