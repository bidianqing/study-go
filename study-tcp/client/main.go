package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败，err:", err)
	}

	fmt.Println("连接成功")

	for {
		fmt.Println("请输入要发送的信息")

		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取终端输入信息失败", err)
		}

		conn.Write([]byte(str))
	}
}
