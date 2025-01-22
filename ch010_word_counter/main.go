package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://www.gutenberg.org/cache/epub/100/pg100.txt
func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordMap := make(map[string]int)
	scanner := bufio.NewScanner(file) // TODO
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			wordMap[strings.ToLower(word)]++
		}
	}

	for word, count := range wordMap {
		fmt.Printf("%s: %d\n", word, count) // TODO
	}
}

// go run main.go | sort -k2 -g

// TODO 改进 标点符号
