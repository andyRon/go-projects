package main

import (
	"encrypt_decrypt/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fileName := "b.zip"
	seed, _ := strconv.ParseInt(os.Args[1], 10, 64)
	encryptFile := "encrypt.enp"
	outputFile := "b.zip"
	if len(os.Args) == 4 {
		encryptFile = os.Args[2]
		outputFile = os.Args[3]
	}

	//fmt.Println("seed=",seed)
	err := utils.DecryptByte(encryptFile, "./file", outputFile, "./file", seed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decrypt the encrypted file \"%s\" to normal file \"%s successfully：\n", encryptFile, outputFile)
}
