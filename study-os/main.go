package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("study-os/user.json")
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	json.Unmarshal(data, &m)

	fmt.Println(m)

	fileInfo, err := os.Stat("study-os/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())
}
