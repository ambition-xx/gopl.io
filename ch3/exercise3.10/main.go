package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buffer bytes.Buffer
	l := len(s)
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))

		if (l - i - 1) % 3 == 0 { //是3的倍数就加一个“，”
			buffer.WriteString(",")
		}
	}
	s = buffer.String()
	return s[:len(s) - 1] // 末尾会多一个逗号,去掉
}
