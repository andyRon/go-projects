package main

import (
	"encrypt_decrypt/utils"
	"os"
	"strconv"
)

// encrypt.exe 123456 a.zip
// ./encrypt 123456 a.zip
// go run encrypt.go 123456 a.zip
func main() {
	secretNum, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fileName := os.Args[2]

	utils.EncryptByte(fileName, "./file", "encrypt.enp", "./file", secretNum)
}
