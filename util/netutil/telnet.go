package netutil

import (
	"fmt"
	"net"
	"time"
)

// CheckPort 检测端口是否可以访问
func Telnet(host string, port, second int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Duration(second)*time.Second)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer conn.Close()
	return true
}
