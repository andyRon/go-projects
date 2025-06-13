package handler

import (
	"encoding/json"
	"fmt"
	"github.com/andyron/go-projects/stl-mongodb/stl"
	"github.com/andyron/go-projects/stl-mongodb/utils"
	"io"
	"log"
	"net/http"
	"time"
)

// 处理STL数据请求

func GetSTLList(w http.ResponseWriter, r *http.Request) {
	// start用于记录处理请求的开始时间
	start := time.Now()
	status := ""
	// 允许跨域查询
	w.Header().Set("Access-Control-Allow-Origin", "*")
	defer func() {
		// 当服务器发生可恢复致命错误时的处理
		if err := recover(); err != nil {
			log.Println("handler \"GetSTLList\" list file error occurred:", err)
			// 返回服务器时间和错误信息
			// Format函数中的"2006-01-02 15:04:05"为Go语言格式化时间所使用的字符串
			status = time.Now().Format("2006-01-02 15:04:05") +
				// 使用fmt.Sprintf拼接字符串
				fmt.Sprintf("get stl file list error:%s", err)
			// 调用标准库的json.Marshal序列化ResponseStatus结构体变量
			response, _ := json.Marshal(ResponseStatus{Status: status})
			_, err = w.Write(response)
			if err != nil {
				log.Println("handler \"GetSTLList\" write response status failed:", err)
				return
			}
		}
	}()
	stlFileList, err := utils.GetFileList("assets/upload")
	if err != nil {
		// 当获取文件列表错误时，打印错误信息并立即返回
		utils.ErrorHandler(err, "handler.GetSTLList list stl file")
		return
	}
	res, err := json.Marshal(STLFileList{STLList: stlFileList})
	if err != nil {
		// 当序列化文件列表信息为字节码发生错误时，打印错误信息并立即返回
		utils.ErrorHandler(err, "handler.GetSTLList marshal STLFileList data")
		return
	}
	_, err = w.Write(res)
	if err != nil {
		// 当给客户端返回响应信息发生错误时，打印错误信息并立即返回
		utils.ErrorHandler(err, "handler.GetSTLList write response")
		return
	}
	// 正确处理完成后，打印处理结束信息，以及处理耗时
	log.Printf("write list STL file response finished successfully, cost time %s\n", -start.Sub(time.Now()))
}

func SaveSTLMongo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status := ""
	// 允许跨域
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// 只接收POST方法
	if r.Method != "POST" {
		status = time.Now().Format("2006-01-02 15:04:05") +
			"server only accept 'POST' method"
		// 调用标准库的json.Marshal序列化ResponseStatus结构体变量
		response, _ := json.Marshal(ResponseStatus{Status: status})
		_, err := w.Write(response)
		if err != nil {
			log.Println("handler \"SaveSTLMongo\" write response status failed:", err)
			return
		}
		return
	}
	defer func() {
		// 当发生可恢复致命错误时的处理
		if err := recover(); err != nil {
			log.Println("handler \"SaveSTLMongo\" list file error occurred:", err)
			// 返回服务器时间和错误信息
			status = time.Now().Format("2006-01-02 15:04:05") +
				//使用fmt.Sprintf拼接字符串
				fmt.Sprintf(", save stl file to MongoDB error:%s", err)
			// 调用标准库的json.Marshal序列化ResponseStatus结构体变量
			response, _ := json.Marshal(ResponseStatus{Status: status})
			_, err = w.Write(response)
			if err != nil {
				log.Println("handler \"SaveSTLMongo\" write response status failed:", err)
				return
			}
		}
	}()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("handler \"SaveSTLMongo\" read request body error:", err)
		return
	}
	defer func() {
		// 处理结束前关闭请求体r.Body
		err := r.Body.Close()
		if err != nil {
			log.Println("handler \"SaveSTLMongo\" close the request body error occurred:", err)
			return
		}
	}()
	stlFileName := STLFileName{}
	if err = json.Unmarshal(body, &stlFileName); err != nil {
		log.Println("handler \"SaveSTLMongo\" unmarshal stlFileName error occurred:", err)
		return
	}
	stlData, err := stl.ReadSTLFile(stlFileName.Name, "assets/upload")
	if err != nil {
		// 当获取stl文件数据发生错误时，打印错误信息并立即返回
		log.Println("handler \"SaveSTLMongo\" read STL file data error occurred:", err)
		return
	}
	err = stl.SaveSTLMongo(stlData,
		"gopher", "gopher2021",
		"127.0.0.1", 27017, "STL", "Binary", 10)
	if err != nil {
		log.Println("handler \"SaveSTLMongo\" save STL file data to MongoDB error occurred:", err)
		return
	}
	status = time.Now().Format("2006-01-02 15:04:05") +
		fmt.Sprintf(", save stl file %s to MongoDB successfully", stlFileName.Name)
	// 调用标准库的json.Marshal序列化ResponseStatus结构体变量
	response, _ := json.Marshal(ResponseStatus{Status: status})
	_, err = w.Write(response)
	if err != nil {
		// 当给客户端返回响应信息发生错误时，打印错误信息并立即返回
		utils.ErrorHandler(err, "handler.GetSTLList write response")
		return
	}
	// 正确处理完成后，打印处理结束信息，以及处理耗时
	log.Printf("write STL file data to MongoDB finished successfully, cost time %s\n", -start.Sub(time.Now()))
}

func QuerySTLMongo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status := ""
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "POST" {
		status = time.Now().Format("2006-01-02 15:04:05") +
			"server only accept 'POST' method"
		// 调用标准库的json.Marshal序列化ResponseStatus结构体变量
		response, _ := json.Marshal(ResponseStatus{Status: status})
		_, err := w.Write(response)
		if err != nil {
			log.Println("handler \"QuerySTLMongo\" write response status failed:", err)
			return
		}
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("handler \"QuerySTLMongo\" read request body error:", err)
		return
	}
	defer func() {
		// 处理结束前关闭请求体r.Body
		err := r.Body.Close()
		if err != nil {
			log.Println("handler \"QuerySTLMongo\" close the request body error occurred:", err)
			return
		}
	}()
	stlFileName := STLFileName{}
	defer func() {
		// 当发生可恢复致命错误时的处理
		if err := recover(); err != nil {
			log.Println("handler \"QuerySTLMongo\" list file error occurred:", err)
			status = time.Now().Format("2006-01-02 15:04:05") +
				fmt.Sprintf(", query stl data(name=%s) in MongoDB error:%s", stlFileName.Name, err)
			response, _ := json.Marshal(ResponseStatus{Status: status})
			_, err = w.Write(response)
			if err != nil {
				log.Println("handler \"QuerySTLMongo\" write response status failed:", err)
				return
			}
		}
	}()
	if err = json.Unmarshal(body, &stlFileName); err != nil {
		log.Println("handler \"QuerySTLMongo\" unmarshal stlFileName error occurred:", err)
		return
	}
	//fmt.Println(stlFileName.Name)
	modelSTL, err := stl.QuerySTLMongo(stlFileName.Name,
		"gopher", "gopher2021",
		"127.0.0.1", 27017, "STL", "Binary", 10)
	if err != nil {
		// 当获取文件列表错误时，打印错误信息并立即返回
		log.Println("handler \"QuerySTLMongo\" read STL file data error occurred:", err)
		return
	}
	response, _ := json.Marshal(modelSTL)
	_, err = w.Write(response)
	if err != nil {
		// 当给客户端返回响应信息发生错误时，打印错误信息并立即返回
		utils.ErrorHandler(err, "handler.GetSTLList write response")
		return
	}
	// 正确处理完成后，打印处理结束信息，以及处理耗时
	log.Printf("write STL file data to MongoDB finished successfully, cost time %s\n", -start.Sub(time.Now()))
}
