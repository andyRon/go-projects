package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		files, _ := os.ReadDir("./")
		for _, file := range files {
			fmt.Print(file.Name(), " ")
		}
	} else {
		files, _ := os.ReadDir(os.Args[1])
		for _, file := range files {
			fmt.Print(file.Name(), " ")
		}
	}
}
