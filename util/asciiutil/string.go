package asciiutil

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"unicode"
)

func Ascii2String(data []byte) string {
	var str string
	for _, b := range data {
		if unicode.IsPrint(int32(b)) {
			//fmt.Printf("%c 是可打印字符\n", char)
			str += fmt.Sprintf("%c", b)
		} else {
			//fmt.Printf("%c 不是可打印字符\n", char)
			str += fmt.Sprintf(".")
		}
	}
	return str
}

func SplitHostPort(address string) (string, uint64, error) {
	if address == "" {
		return "", 0, errors.New("address is nil")
	}
	host, port, err := net.SplitHostPort(address)
	if err == nil {
		num, err1 := strconv.Atoi(port)
		if err1 == nil {
			return host, uint64(num), err
		} else {
			return host, uint64(num), err1
		}
	} else {
		return host, 0, err
	}
}
