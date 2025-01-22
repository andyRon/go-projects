package main

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// TODO
func main() {
	fileHaslList := make(map[string][]byte)
	archive_flag := false
	for {
		filepath.WalkDir("input", func(path string, info os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			// Do we have this file
			if hash, ok := fileHaslList[path]; ok {
				// take the hash of this file
				file, _ := os.Open(path)
				h := md5.New()
				io.Copy(h, file)
				nhash := h.Sum(nil)
				if !bytes.Equal(hash, nhash) {
					archive_flag = true
					return errors.New("Rearchive")
				}
				file.Close()
			} else {
				archive_flag = true
				return errors.New("Rearchive")
			}
			return nil
		})

		if archive_flag {
			// create the archive
			os.Remove("output.zip")
			outfile, _ := os.Create("output.zip")
			w := zip.NewWriter(outfile)
			log.Println("Creating a new archive")
			filepath.WalkDir("input", func(path string, info os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				// creatr a hash of the new file
				file, _ := os.Open(path)
				h := md5.New()
				io.Copy(h, file)
				nhash := h.Sum(nil)
				fileHaslList[path] = nhash

				// compress the file
				f, _ := w.Create(path)
				file.Seek(0, io.SeekStart)
				io.Copy(f, file)
				file.Close()
				return nil
			})
			archive_flag = false
			w.Close()
		}
		time.Sleep(time.Second * 1)
	}

}

/*
res: https://www.youtube.com/watch?v=xbI2ELnVGAo&list=PLYmIsLVSssdIxboWkccLoRWdX32-tpg21&index=7
*/
