package utils

import (
	"path/filepath"
	"runtime"
)

func GetPackageInfo() (pkgName, fileName, funcName string) {
	// 获取调用者信息
	pc, file, _, _ := runtime.Caller(1)
	// 解析包名
	dir := filepath.Dir(file)
	pkgName = filepath.Base(dir)
	// 文件名
	fileName = filepath.Base(file)
	// 函数名
	funcName = runtime.FuncForPC(pc).Name()
	return
}
