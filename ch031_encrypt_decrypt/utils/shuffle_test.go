package utils

import (
	"fmt"
	"testing"
)

func TestShuffleN(t *testing.T) {
	res := ShuffleN(10, 10088)
	t.Log(res)
}

func TestShuffle(t *testing.T) {
	l := 10
	origin := make([]int, l, l)
	for i, _ := range origin {
		origin[i] = i
	}
	t.Log("origin=", origin)
	res := Shuffle(origin, 666666)
	t.Log("res=", res)
}

func TestGenByteMap(t *testing.T) {
	m := GenByteMap(10088)
	t.Log("m=", m)
	n := make(map[uint8]uint8, 256)
	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
		n[v] = k
	}
	t.Log("n=", n)
}
