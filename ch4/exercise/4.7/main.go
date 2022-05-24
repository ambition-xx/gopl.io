// 练习 4.7: 修改 reverse 函数用于原地反转 UTF-8 编码的[]byte。是否可以不用分配额外的内存?
package main
import (
	"fmt"
)

func main() {
	s := []byte("abc def   ghi jk")
	reverse(s)
	fmt.Println(string(s))
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}