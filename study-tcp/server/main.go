package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败，err:", err)
		return
	}

	fmt.Println("监听成功")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept失败，err:", err)
			return
		} else {
			fmt.Println("Accept成功 客户端信息", conn.RemoteAddr().String())
		}

		go func(c net.Conn) {
			defer c.Close()

			for {
				b := make([]byte, 1024)
				n, err := c.Read(b)
				if err != nil {
					return
				}

				fmt.Println(string(b[0:n]))
			}

		}(conn)
	}
}
