package main

import (
	"github.com/andyron/panel/define"
	"syscall"
)

func main() {
	define.PID = syscall.Getpid()
	// TODO
}
