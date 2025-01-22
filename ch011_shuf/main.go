package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	for _, line := range lines {
		fmt.Println(line)
	}
}

// ls ~/myfield/tmp | go run main.go
