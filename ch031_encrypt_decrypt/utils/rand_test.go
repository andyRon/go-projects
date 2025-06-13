package utils

import "testing"

func TestGenRand(t *testing.T) {
	r, err := GenRand(256)
	if err != nil {
		t.Log(err)
	}
	t.Log("r=", r)
}

func TestGenConfuseBytes(t *testing.T) {
	cb, err := GenConfuseBytes(10)
	if err != nil {
		t.Log("error:", err)
		return
	}
	t.Log(cb)
}
