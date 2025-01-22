package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	filepath.WalkDir("./collection", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".jpg", ".png", ".mp3", ".mp4":
			os.Rename(path, "./collection/media/"+filepath.Base(path))
		case ".pdf", ".doc", ".md", ".html":
			os.Rename(path, "./collection/docs/"+filepath.Base(path))
		}
		return err
	})
}
