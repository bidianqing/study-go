package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(time.Now().AddDate(0, 0, 4))
}
