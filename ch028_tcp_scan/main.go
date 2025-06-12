package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// 配置命令行参数
	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port on which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout for each port")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{} // TODO
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)

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

// ./main -hostname baidu.com -start-port 80 -end-port 200 -timeout 100ms
