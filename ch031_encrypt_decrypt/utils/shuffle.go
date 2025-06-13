package utils

import "math/rand"

// 用于开发构建字节映射变换的函数

// 设定一个固定的偏移量，根据需要可修改此数值
const seedOffset int64 = 10077

// ShuffleN 输入一个n值，返回[0,n)区间数值的一个伪随机排列
func ShuffleN(n int, seed int64) (res []int) {
	randSource := rand.New(rand.NewSource(seed))
	res = randSource.Perm(n)
	return
}

// Shuffle 输入一个int类型的切片，将切片内的元素进行随机的重新排列返回新的切片
func Shuffle(origin []int, seed int64) (res []int) {
	randSource := rand.New(rand.NewSource(seed))
	l := len(origin)
	res = make([]int, l, l)
	perm := randSource.Perm(l)
	for i, randIndex := range perm {
		res[i] = origin[randIndex]
	}
	return
}

func GenByteMap(seed int64) (m map[byte]byte) {
	m = make(map[byte]byte, 256)
	origin := make([]byte, 256, 256)
	for i, _ := range origin {
		origin[i] = byte(i)
	}

	// 将0~127范围内的数进行重新排列，得到一个切片
	permTop := rand.New(rand.NewSource(seed)).Perm(128)
	for i, _ := range permTop {
		// 将切片中的每一个值增加128的偏移量，进而每个数值变换为128~255范围，作为原始字节数值低128位映射值
		permTop[i] += 128
	}
	// 得到原始字节的高128位映射值，取值范围为0~127
	permTail := rand.New(rand.NewSource(seed + seedOffset)).Perm(128)
	// 合并映射后的两个切片，作为一个新的切片
	perm := append(permTop, permTail...)
	for i, v := range perm {
		m[origin[i]] = byte(v)
	}
	return
}

// ReverseByteMap 反转map类型的键key和值value
func ReverseByteMap(m map[byte]byte) (rm map[byte]byte) {
	rm = make(map[byte]byte, len(m))
	for k, v := range m {
		rm[v] = k
	}
	return
}
