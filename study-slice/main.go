package main

/*
https://github.com/golang/go/blob/master/src/runtime/slice.go#L15

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

无论是 var 声明定义的 slice 变量，还是 make(xxx，num) 创建的 slice 变量，slice 管理结构是已经分配出来了的（也就是 struct slice 结构 ）。
所以， 对于 slice 来说，其实并不需要 make 创建的才能使用，直接用 var 定义出来的 slice 也能直接使用

make时，如果传一个数字n1，那么切片的长度和容量都是n1，如果传了两个数字n1，n2 则n1是长度，n2是容量，n1不能大于n2

切片的扩容机制很复杂，每个版本都有所不同，跟机器是32位和64位也有关系
*/

func main() {
	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := s1[1:]

	// s1[1] = 100
	// fmt.Println(s1)
	// fmt.Println(s2)

	// 发生扩容，append操作的会复制一个新的数组
	// s1 := []int{11, 22, 33}
	// s2 := append(s1, 44) // 发生了扩容 s2的长度是4，容量是6
	// s1[0] = 999

	// s3 := append(s2, 55) // 给s2追加了55，没有发生扩容，所以s2和s3指向的是同一个底层数组
	// s2[3] = 100
	// fmt.Println(s1) //[999 22 33]
	// fmt.Println(s2) //[11 22 33 100]
	// fmt.Println(s3) //[11 22 33 100 55]

	// 没有发生扩容，append操作的还是同一个底层数组
	// s1 := make([]int, 1, 3)
	// s2 := append(s1, 1, 2)
	// s1[0] = 999
	// fmt.Println(s1) // [999]
	// fmt.Println(s2) //[999 1 2]
}

type User struct {
	Id   int
	Name string
}
