package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	ports := []int{}

	wg := &sync.WaitGroup{}
	timeout := time.Millisecond * 200
	for port := 1; port < 100; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen("baidu.com", p, timeout)
			if opened {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Sprintf("Open ports: %v\n", ports)

}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}
