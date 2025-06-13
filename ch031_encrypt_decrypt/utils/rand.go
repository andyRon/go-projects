package utils

import (
	crand "crypto/rand"
	"log"
	"math/big"
)

// 用于开发生成混淆字节的函数

// GenRand 产生随机字节
// 将输入的整数作为随机数（正整数）种子，返回一个随机的byte类型值
func GenRand(maxInt int64) (b byte, err error) {
	r, err := crand.Int(crand.Reader, big.NewInt(maxInt))
	b = byte(r.Int64())
	return
}

// GenConfuseBytes 根据需要的字节长度，生成随机byte数值构成的切片，用于文件混淆
func GenConfuseBytes(n uint) (cb []byte, err error) {
	cb = make([]byte, n, n)
	for i, _ := range cb {
		b, err2 := GenRand(256)
		if err2 != nil {
			log.Println("utils/rand.go/GenConfuseBytes:generate random bytes error:", err)
			err = err2
			return
		}
		cb[i] = b
	}
	return
}
