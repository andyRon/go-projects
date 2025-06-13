package main

import (
	"fmt"
	"github.com/andyron/go-projects/stl-mongodb/handler"
	"github.com/andyron/go-projects/stl-mongodb/utils"
	"net/http"
	"time"
)

func main() {
	// 创建一个名为mux的http multiplexer(多路复用器)
	// 不使用http.DefaultServerMux
	mux := http.NewServeMux()

	// 定义服务器的IP地址
	// 当使用“127.0.0.1”时，服务器程序只能被本机访问。
	// 当使用“0.0.0.0”时，服务器程序可被所在网络上的任意计算机访问
	ip := "0.0.0.0"
	//定义服务器程序的运行端口，端口范围“0~65535”，建议使用大于8000的非活动端口
	port := "8888"

	//当用户访问"127.0.0.1:8888/"或"127.0.0.1:8888/ping"两个路由地址时，调用handler包中的Ping函数进行处理
	mux.HandleFunc("/", handler.Ping)
	mux.HandleFunc("/ping", handler.Ping)
	//当用户访问"127.0.0.1:8888/get-stl-list"时，调用handler包中的GetSTLList函数进行处理
	mux.HandleFunc("/get-stl-list", handler.GetSTLList)
	//当用户访问"127.0.0.1:8888/save-stl-mongo"时，调用handler包中的SaveSTLMongo函数进行处理
	mux.HandleFunc("/save-stl-mongo", handler.SaveSTLMongo)
	//当用户访问"127.0.0.1:8888/query-stl-mongo"时，调用handler包中的QuerySTLMongo函数进行处理
	mux.HandleFunc("/query-stl-mongo", handler.QuerySTLMongo)

	//使用Go语言标准库http创建Server
	server := &http.Server{
		Addr:              ip + ":" + port, //服务器程序运行的ip地址和端口
		Handler:           mux,             //服务器程序所使用的http multiplexer
		TLSConfig:         nil,
		ReadTimeout:       time.Duration(150 * int64(time.Second)), //定义读超时为150秒，读者可灵活调整
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Duration(600 * int64(time.Second)), //定义写超时为150秒，读者可灵活调整
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	fmt.Println("server started at " + ip + ":" + port)
	err := server.ListenAndServe()
	if err != nil {
		//定义一个字符串变量存储错误信息
		waitStr := ""
		//当server对象在调用ListenAndServe方法后发生错误时，此时应强制退出程序
		//使用fmt.Scanf函数会将程序运行过程卡住，当按下回车键后，程序继续执行。
		//主要目的是防止程序退出前命令行窗口一闪而过。
		_, _ = fmt.Scanf("panic error %s occurred, press \"Enter\" key to exit...\n", &waitStr)
		//使用utils包中定义的panic错误的处理函数，程序强制退出
		utils.PanicHandler(err, "main.go")
	}
}
