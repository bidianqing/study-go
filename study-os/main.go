package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"gopkg.in/ini.v1"
)

func main() {
	// data, err := os.ReadFile("study-os/user.json")
	// if err != nil {
	// 	panic(err)
	// }

	// m := make(map[string]interface{})
	// json.Unmarshal(data, &m)

	// fmt.Println(m)

	// fileInfo, err := os.Stat("study-os/a.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(fileInfo.Name())

	envs := os.Environ()
	for _, v := range envs {
		fmt.Println(v)
	}

	data, err := os.ReadFile("study-os/frpc.ini")
	if err != nil {
		panic(err)
	}
	file, err := ini.Load(data)
	if err != nil {
		panic(err)
	}

	_, err = file.GetSection("common")
	if err != nil {
		panic(err)
	}

	Command("dotnet", "--info")
}

func Command(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	resp, _ := io.ReadAll(stdout)
	fmt.Println(string(resp))
}
