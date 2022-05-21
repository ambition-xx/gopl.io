package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buffer bytes.Buffer
	var symbol string
	if s[0] == '+' || s[0] == '-'{
		symbol = string(s[0])
		s = s[1:]
		buffer.WriteString(symbol)
	} 
	arr := strings.Split(s, ".")
	length := len(arr[0])
	for i := 0; i < length; i++ {
		buffer.WriteString(string(s[i]))

		if length != i + 1 && (length-i-1)%3 == 0 { //是3的倍数就加一个“，”
			buffer.WriteString(",")
		}
	}
	if len(arr) > 1{
		buffer.WriteString(".")
		buffer.WriteString(arr[1])
	}
	s = buffer.String()
	return s // 末尾会多一个逗号,去掉
}
