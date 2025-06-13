package handler

import (
	"fmt"
	"github.com/andyron/go-projects/stl-mongodb/utils"
	"log"
	"net/http"
	"time"
)

// 用于测试服务器程序连通性，给客户端返回一个字符串“Pong”+服务器时间
// Go语言中作为http服务器的处理函数，其形参必须为http.ResponseWriter和*http.Request，且顺序不可颠倒

func Ping(w http.ResponseWriter, r *http.Request) {
	// 当发生致命错误时，调用utils包中的RecoverHandler进行故障恢复，防止意外退出
	defer utils.RecoverHandler("handler.Ping")
	defer func() {
		//关闭客户端的http请求体r.Body
		err := r.Body.Close()
		if err != nil {
			//符号\"用于转义，目的是在字符串中输出双引号"
			log.Println("handler \"Ping\" close the request body error occurred:", err)
			return
		}
	}()

	// 查看客户端请求的方法
	fmt.Println("handler \"Ping\": client request method is", r.Method)
	// w即为给客户端返回内容的ResponseWriter，字面意思可以理解为响应写入器
	// w返回的数据必须为字节切片格式，所以使用[]byte()对返回的字符串进行强制的数据类型转换
	num, err := w.Write([]byte("Pong!" + time.Now().Format("2006-01-02 15:04:05")))
	// 如果w在返回响应时产生错误，则调用utils包的一般错误处理函数进行处理
	utils.ErrorHandler(err, "handler.Ping write response")
	// 在服务器段输出返回给客户端的字节数
	log.Println("Ping handler write", num, "bytes")
}
