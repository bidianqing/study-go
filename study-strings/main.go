package main

import (
	"fmt"
	"strings"
)

func main() {
	// 去除字符长前面和后面的空格
	name := "    tom        "
	fmt.Println(strings.Trim(name, " "))
	fmt.Println(strings.TrimSpace(name))

	// 判断是否以指定字符串开头并去除
	token := "Bearer 123456"
	fmt.Println(strings.HasPrefix(token, "Bearer "))
	fmt.Println(strings.TrimPrefix(token, "Bearer "))
}
