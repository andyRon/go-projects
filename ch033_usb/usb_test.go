package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestFindUSBDrive(t *testing.T) {
	usbDrive, err := findUSBDrive()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Use Drive found at:", usbDrive)
}

func TestReadFileFromUSB(t *testing.T) {
	data, err := readFileFromUSB(filepath.Join(drives[0], "example.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file content:", string(data))
}

func TestGetFileStat(t *testing.T) {
	fileStat, err := getFileStat(filepath.Join(drives[0], "example.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("File Name: %s\n", fileStat.Name())
	fmt.Printf("File Size: %d bytes\n", fileStat.Size())
	fmt.Printf("Modified Time: %s\n", fileStat.ModTime().Format(time.DateTime))
}

func TestWriteFileToUSB(t *testing.T) {
	err := writeFileToUSB(filepath.Join(drives[0], "output.txt"), []byte("hello USB!"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write file to usb success")
}

func TestReadFileWithBufio(t *testing.T) {
	data, err := readFileFromUSB(filepath.Join(drives[0], "example.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file content:", string(data))
}

func TestProcessFile(t *testing.T) {
	entries, err := os.ReadDir(drives[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := filepath.Join(drives[0], entry.Name())
			wg.Add(1) // TODO
			go processFile(filePath, &wg)
		}
	}
	wg.Wait()
}
