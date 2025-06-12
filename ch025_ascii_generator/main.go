package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"syscall"
)

func main() {
	//l1()
	//l2()
	//l3()
	//l4()

	// 验证文件内容
	//verifyFile("ascii1.txt")
	//verifyFile("ascii2.txt")
	//verifyFile("ascii3.txt")
	//verifyFile("ascii4.txt")
}

// 普通生成
func l1() {
	file, err := os.Create("ascii1.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	for i := 0; i < 128; i++ {
		file.WriteString(fmt.Sprintf("%c\n", i))
	}

	fmt.Println("ASCII file created successfully.")
}

// 带缓冲的生成
func l2() {
	file, err := os.Create("ascii2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 128; i++ {
		writer.WriteString(fmt.Sprintf("%c\n", i))
	}
	writer.Flush()

	fmt.Println("ASCII file generated successfully with buffered writer.")
}

// 并发写入生成
func l3() {
	file, err := os.Create("ascii3.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	partSize := 32 // 将任务分成4部分

	for i := 0; i < 128; i += partSize {
		wg.Add(1)
		go writeASCIIChars(file, i, i+partSize, &wg)
	}

	wg.Wait()
	fmt.Println("ASCII file generated successfully with concurrent writing.")
}

func writeASCIIChars(file *os.File, start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	writer := bufio.NewWriter(file)
	for i := start; i < end; i++ {
		writer.WriteString(fmt.Sprintf("%c\n", i))
	}
	writer.Flush()
}

// 使用内存映射文件（Memory-Mapped Files）来提高性能
func l4() {
	file, err := os.Create("ascii4.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 设置文件大小
	if err := file.Truncate(128); err != nil {
		fmt.Println("Error setting file size:", err)
		return
	}

	// 内存映射文件
	data, err := syscall.Mmap(int(file.Fd()), 0, 128, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("Error memory-mapping file:", err)
		return
	}
	defer syscall.Munmap(data)

	// 写入ASCII字符
	for i := 0; i < 128; i++ {
		data[i] = byte(i)
	}

	fmt.Println("ASCII file generated successfully with memory-mapped file.")
}

// 验证文件内容
func verifyFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Text() != string(rune(i)) {
			fmt.Printf("Mismatch at character %d\n", i)
			return
		}
	}
	fmt.Println("File verification successful.")
}
