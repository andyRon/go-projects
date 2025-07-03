package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// go处理外部存储设备（如U盘）的数据和普通目录一样，只需要找到对应挂载位置。
// 通常情况下，Mac OS 外部存储设备会自动挂载到 /Volumes 目录下。
// 在Linux系统中，U盘通常挂载在/media或/mnt目录下；在Windows系统中，U盘则可能出现在D:、E:等盘符中。

var drives = []string{"/Volumes/FX-SSD-PS2000/test"}

func findUSBDrive() ([]string, error) {
	var res []string
	drives := []string{"/Volumes/"}
	for _, drive := range drives {
		entries, err := os.ReadDir(drive)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				res = append(res, filepath.Join(drive, entry.Name()))
			}
		}
	}
	if len(res) > 0 {
		return res, nil
	} else {
		return res, fmt.Errorf("USB drive not found")
	}
}

// 使用os包中的Open和Read函数可以读取U盘中的文件。
func readFileFromUSB(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]byte, 1024)
	var bytesRead []byte
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		bytesRead = append(bytesRead, data[:n]...)
	}
	return bytesRead, nil
}

// 获取文件属性  使用os包中的Stat函数可以获取文件的属性信息，如文件大小、修改时间等。
func getFileStat(filePath string) (os.FileInfo, error) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	return fileStat, nil
}

// 文件写入操作 使用os包中的OpenFile和Write函数可以实现对文件的写入操作。
func writeFileToUSB(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// bufio包提供了一个缓冲区，可以显著提高文件读写的效率
func readFileWithBufio(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var bytesRead []byte
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		bytesRead = append(bytesRead, buffer[:n]...)
	}
	return bytesRead, nil
}

// 并发处理文件
func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("File: %s, Size: %d bytes\n", fileStat.Name(), fileStat.Size())
}
