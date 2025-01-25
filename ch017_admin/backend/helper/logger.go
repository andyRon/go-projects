package helper

import (
	"fmt"
	. "github.com/andyron/mini-admin/define"
	"runtime"
	"strconv"
	"time"
)

func dataFormat(showDetail bool, format string, v ...interface{}) string {
	prefix := "[" + FrameName + "] " + time.Now().Format(DateTimeLayout) + " "
	if showDetail {
		_, file, line, _ := runtime.Caller(2)
		prefix += "file: " + file + " line: " + strconv.Itoa(line) + " ==> "
	}
	return prefix + fmt.Sprintf(format, v...)
}

func Info(format string, v ...interface{}) {
	fmt.Printf("\033[32m%s\033[0m\n", dataFormat(false, format, v...))
}

func Error(format string, v ...interface{}) {
	fmt.Printf("\033[31m%s\033[0m\n", dataFormat(true, format, v...))
}
