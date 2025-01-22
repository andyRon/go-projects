package main

import (
	"io"
	"os"
)

// Linux命令 cat

func main() {
	for _, arg := range os.Args[1:] {
		file, _ := os.Open(arg)
		io.Copy(os.Stdout, file)
		file.Close()
	}
}
