package main

import (
	"fmt"
	"os"
)

// go run main.go
// 清空文件的三种方法

func main() {
	//truncate()

	//writeFile()

	create()
}

// 1️⃣ os.Truncate函数可以将文件截断到指定的大小，如果大小为0，则相当于清空文件。
func truncate() {
	filePath := "./test.txt"
	err := truncateFile(filePath)
	if err != nil {
		fmt.Println("Error truncating file:", err)
	} else {
		fmt.Println("File truncated successfully")
	}
}

func truncateFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		return err
	}
	return nil
}

// 2️⃣ 通过os.WriteFile函数将空字符串写入文件，也可以实现清空文件的效果。
func writeFile() {
	filePath := "test.txt"
	err := clearFile(filePath)
	if err != nil {
		fmt.Println("Error clearing file:", err)
	} else {
		fmt.Println("File cleared successfully")
	}
}

func clearFile(filePath string) error {
	return os.WriteFile(filePath, []byte(""), 0666)
}

// 3️⃣ 通过os.Create函数创建一个新的文件，也可以实现清空文件的效果。
func create() {
	filePath := "test.txt"
	err := recreateFile(filePath)
	if err != nil {
		fmt.Println("Error recreating file:", err)
	} else {
		fmt.Println("File recreated successfully")
	}
}

func recreateFile(filePath string) error {
	_, err := os.Create(filePath)
	return err
}
