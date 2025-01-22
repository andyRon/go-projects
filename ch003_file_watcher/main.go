package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileInfo, err := os.Stat("index.html")
	if err != nil {
		panic(err)
	}
	oldTime := fileInfo.ModTime()

	for {
		fileInfo, err := os.Stat("index.html")
		if err != nil {
			panic(err)
		}
		newTime := fileInfo.ModTime()

		if newTime.After(oldTime) {
			fmt.Println("file has been modified in ", newTime)
			oldTime = newTime
		}
		time.Sleep(200 * time.Millisecond)
	}
}
