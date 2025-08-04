package main

import (
	"io"
	"os"
)

// 实现命令cat

func main() {
	for _, arg := range os.Args[1:] {
		file, _ := os.Open(arg)
		io.Copy(os.Stdout, file)
		file.Close()
	}
}
