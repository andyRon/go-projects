package utils

import (
	"log"
	"os"
	"path/filepath"
)

// 编写加密字节和解密字节的核心函数

// EncryptByte 加密文件
// seed 即随机数种子，也就是程序在运行时读取命令行用户输入的参数
func EncryptByte(inputFile, inputPath, outputFile, outputPath string, seed int64) (err error) {
	fileBytes, err := os.ReadFile(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte:read input file error:", err)
		return
	}
	// 构造加密文件内容的字节切片
	outputFileBytes := make([]byte, 0)
	m := GenByteMap(seed)
	for i, _ := range fileBytes {
		outputFileBytes = append(outputFileBytes, m[fileBytes[i]])
	}
	// 生成放在文件头部的混淆字节码
	confuseBytesTop, err := GenConfuseBytes(77)
	if err != nil {
		log.Println("utils/encrypt.go/EncryptByte:gen top confuse bytes error:", err)
		return
	}
	// 生成放在文件尾部的混淆字节码
	confuseBytesTail, err := GenConfuseBytes(66)
	if err != nil {
		log.Println("utils/encrypt.go/EncryptByte:gen tail confuse bytes error:", err)
		return
	}
	// 合并头部的混淆字节切片和加密后的字节切片
	outputFileBytes = append(confuseBytesTop, outputFileBytes...)
	// 合并加密后的字节切片和尾部的混淆字节码切片
	outputFileBytes = append(outputFileBytes, confuseBytesTail...)
	// 将合并后的字节切片输出到文件
	outputFP, err := os.Create(filepath.Join(outputPath, outputFile))
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte,create output file error:", err)
		return
	}

	_, err = outputFP.Write(outputFileBytes)
	if err != nil {
		log.Println("utils/enigma.go/EncryptByte,write output file error:", err)
		return
	}

	// 计算原始文件的散列值并保存
	hashFileName := "sha256_origin.txt"
	CalcSha256(inputFile, inputPath, outputPath, hashFileName)
	return
}

// DecryptByte 解密文件
func DecryptByte(inputFile, inputPath, outputFile, outputPath string, seed int64) (err error) {
	fileBytes, err := os.ReadFile(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte,read input file error:", err)
	}

	l := len(fileBytes)
	// 首先去掉头部和尾部用于混淆的字节
	fileBytesCore := fileBytes[77 : l-66]
	// 初始化一个用于存放解密文件字节的切片
	outputFileBytes := make([]byte, l-77-66, l-77-66)
	// 根据密码得到字节码对应表
	m := GenByteMap(seed)

	// 反转map的key和value
	n := ReverseByteMap(m)
	// 逐个字节的还原原始的文件字节内容
	for i, _ := range fileBytesCore {
		outputFileBytes[i] = n[fileBytesCore[i]]
	}
	// 将还原后的字节切片输出到文件
	outputFP, err := os.Create(filepath.Join(outputPath, outputFile))
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte,create output file error:", err)
		return
	}
	_, err = outputFP.Write(outputFileBytes)
	if err != nil {
		log.Println("utils/enigma.go/DecryptByte,write output file error:", err)
		return
	}
	hashFileName := "sha256_decode.txt"
	// 计算解密后的散列值用于和原始文件的散列值进行比较
	CalcSha256(outputFile, outputPath, outputPath, hashFileName)
	return
}
