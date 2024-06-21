package asciiutil

import (
	"fmt"
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
