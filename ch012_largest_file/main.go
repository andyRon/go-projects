package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// Find the largest file in directory

func main() {
	findLargestFile()
}

func findLargestFile() {
	var pathLF string
	var sizeLF int64

	filepath.WalkDir("../", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}
			if fileInfo.Size() > sizeLF {
				pathLF = path
				sizeLF = fileInfo.Size()
			}
		}
		return nil
	})

	fmt.Println("Largest file:", pathLF, " Size:", sizeLF)
}
