package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 对文件操作的函数集

// ExecutableDir 存储整个项目编译后所在文件夹的绝对路径
var ExecutableDir string

// init函数会运行在utils包中的其他函数运行之前
// 当其他包引入utils包后，utils包中的init()函数会先于其他函数运行
func init() {
	InitExecutableDir()
}

// InitExecutableDir 获取程序编译后的可执行文件所在文件夹名称
func InitExecutableDir() {
	// 获取程序编译后的可执行文件路径executableFilePath，包含可执行文件的文件名
	executableFilePath, err := os.Executable()
	if err != nil {
		log.Println("\"GetFileList\" get executive file dir path error:", err)
		return
	}
	// 先使用Go语言标准库中的filepath.Dir()函数去除可执行文件的文件名
	// 再使用Go语言标准库中的filepath.EvalSymlinks()函数去除符号链接
	// executableDir变量存储的就是dirPath所在的绝对路径
	ExecutableDir, err = filepath.EvalSymlinks(filepath.Dir(executableFilePath))
	if err != nil {
		fmt.Println("eval symlinks error:", err)
		PanicHandler(err, "utils/GetFileList")
	}
}

// GetFileList 传入一个相对于程序可执行文件的相对目录，返回目录下的所有stl文件名的字符串切片
func GetFileList(dirPath string) (fileList []string, err error) {
	files, err := os.ReadDir(filepath.Join(ExecutableDir, dirPath))
	if err != nil {
		log.Println("\"GetFileList\" read dir of \"", dirPath, "\" error:", err)
		return
	}
	fileListLen := len(files)
	fileList = make([]string, fileListLen, fileListLen)
	// 初始化一个uint32变量numFileSTL，用于存放确切的STL文件的数量
	var numFileSTL uint32 = 0
	for _, f := range files {
		if !f.IsDir() {
			//如果f为文件
			//获取文件全名中的后四位，即文件的扩展名，存储在变量extName中
			extName := f.Name()[len(f.Name())-4 : len(f.Name())]
			//如果文件的扩展名extName为".STL"或".stl"
			isSTL := (extName == ".STL") || (extName == ".stl")
			if isSTL {
				//则numFileSTL加1
				numFileSTL++
				//将stl的文件名放入fileList
				fileList[numFileSTL-1] = f.Name()
			}
		}
	}
	fileList = fileList[:numFileSTL]
	return
}
