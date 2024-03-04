package main

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
