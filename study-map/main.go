package main

import (
	"fmt"
)

// 不能使用cap获取map的容量
// key必须是为可比较的类型，不能用做键的数据类型有切片、map和函数
//
//
//
//
//
//

func main() {
	// 只声明，是不能给键赋值的
	var user map[string]string
	fmt.Println(user == nil)
	user = make(map[string]string, 1)
	user["name"] = "tom"
	user["address"] = "北京"

	val, ok := user["age"]
	fmt.Println(val, ok)

	// user := make(map[string]string, 10) // 10为容量
	// user["name"] = "tom"

	// fmt.Println(user)
}
