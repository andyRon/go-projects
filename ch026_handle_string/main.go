package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func main() {
	l1()
	l2()
	l3()

	app1()
	app2()

	yh()
}

// 1️⃣使用[]rune类型处理Unicode字符串
func l1() {
	s := "你好，世界"
	runes := []rune(s)
	fmt.Println(string(runes[0:2])) // 输出: 你好
}

// 2️⃣使用strings包。strings包提供了丰富的字符串处理函数
func l2() {
	s := "www.andyron.top"
	fmt.Println(strings.Split(s, ".")[1])
}

// 3️⃣使用正则表达式
func l3() {
	s := "ID:9527, Name:Andy Ron"
	re := regexp.MustCompile(`ID:(\d+), Name:(\w+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 0 {
		fmt.Println("ID:", matches[1])
		fmt.Println("Name:", matches[2])
	}
}

/* 应用场景 */
// 1️⃣日志分析。在日志分析中，经常需要从日志字符串中提取特定的信息，如时间戳、错误码等。
func app1() {
	log := "2025-01-01 12:00:00 [ERROR] Error code: 404"
	re := regexp.MustCompile(`\[(.*?)\]`)
	level := re.FindString(log)
	fmt.Println("Log Level:", strings.Trim(level, "[]")) // 输出: Log Level: ERROR

	codeRe := regexp.MustCompile(`Error code: (\d+)`)
	codeMatches := codeRe.FindStringSubmatch(log)
	if len(codeMatches) > 0 {
		fmt.Println("Error Code:", codeMatches[1]) // 输出: Error Code: 404
	}
}

// 2️⃣JSON解析。在处理JSON数据时，经常需要将JSON字符串转换为Go语言的数据结构。
func app2() {
	data := `{"name":"Andy Ron","age":18}`
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	fmt.Println("Name:", result["name"])
	fmt.Println("Age:", result["age"])
}

/* 性能优化 */
/*
在处理大量字符串数据时，性能优化尤为重要。以下是一些优化建议：

1 避免重复编译正则表达式：将编译后的正则表达式对象缓存复用。
2 减少字符串拷贝：尽量使用切片而非子字符串创建新的字符串。
3 并行处理：对于大规模数据处理，利用Golang的并发特性进行并行处理。
*/
func yh() {
	logs := []string{
		"2025-01-01 12:00:00 [ERROR] Error code: 404",
		"2025-01-01 12:01:00 [INFO] User logged in",
	}
	re := regexp.MustCompile(`\[(.*?)\]`)
	var wg sync.WaitGroup
	for _, log := range logs {
		wg.Add(1)
		go func(log string) {
			defer wg.Done()
			level := re.FindString(log)
			fmt.Println("Log Level:", strings.Trim(level, "[]"))
		}(log)
	}
	wg.Wait()
}
