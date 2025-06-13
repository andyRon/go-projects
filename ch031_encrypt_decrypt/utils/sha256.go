package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// 用于开发计算文件散列值的函数

// CalcSha256 Go语言标准库中的hash包，通过io包的Copy（）函数将读取的文件字节像水流一样传递给hashSha256变量，
// 并利用格式化字符串的“%x”占位符输出以十六进制数表示的散列值。
// 同时，将散列值保存在文件中，以便授权的解密用户用来比对文件解密后的散列值是否正确。
func CalcSha256(inputFile, inputPath, outputPath, hashFileName string) (err error) {
	fp, err := os.Open(filepath.Join(inputPath, inputFile))
	if err != nil {
		log.Println("utils/sha256.go/CalcSha256,while check hash sha256 read input file error:", err)
		return
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Println("utils/sha256.go/CalcSha256,while check hash sha256 close input file error:", err)
		}
	}()

	hashSha256 := sha256.New()
	if _, err = io.Copy(hashSha256, fp); err != nil {
		log.Println("utils/sha256.go/CalcSha256,while check hash sha256 copy input file error:", err)
		return
	}

	hashSha256String := fmt.Sprintf("%x", hashSha256.Sum(nil))
	fmt.Println("input file hash sha256 code is:", hashSha256String)

	// 将计算得到的文件散列值写入到一个文件中并保存
	hashCodeFileFP, err := os.Create(filepath.Join(outputPath, hashFileName))
	if err != nil {
		log.Println("utils/sha256.go/CalcSha256,create hash file error:", err)
		return
	}

	fileContent := "file name:" + inputFile + "\r\n" + "sha256 code:" + hashSha256String + "\r\n"
	_, err = hashCodeFileFP.Write([]byte(fileContent))
	if err != nil {
		log.Println("utils/sha256.go/CalcSha256,write hash file error:", err)
		return
	}
	defer func() {
		if err = hashCodeFileFP.Close(); err != nil {
			log.Println("utils/sha256.go/CalcSha256,close hash file error:", err)
			return
		}
	}()
	return
}
