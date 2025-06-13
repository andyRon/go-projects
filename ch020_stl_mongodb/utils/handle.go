package utils

import "log"

// 错误处理函数

// RecoverHandler 可恢复的错误处理函数，阻止程序退出。
// 用于虽然发生致命错误，但不让整个程序退出的情景。
func RecoverHandler(info string) {
	if err := recover(); err != nil {
		log.Println("\""+info+"\", recover error occurred:", err)
	}
}

// ErrorHandler 用于一般的错误处理
func ErrorHandler(err error, info string) {
	if err != nil {
		log.Println("\""+info+"\", error occurred:", err)
	}
}

// PanicHandler 当发生导致整个程序不能再继续正常工作的严重错误时(如启动http服务器失败)，强制退出程序
func PanicHandler(err error, info string) {
	if err != nil {
		// log.Fatalln用于打印错误信息并强制退出程序
		log.Fatalln(info, " exited, panic error occurred:", err)
	}
}
