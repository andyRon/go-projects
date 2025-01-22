package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	extCount := make(map[string]int)

	filepath.WalkDir("/Users/andyron/", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		ext := filepath.Ext(path)
		extCount[ext]++
		return err
	})

	for e, c := range extCount {
		fmt.Println(e, c)
	}
}
