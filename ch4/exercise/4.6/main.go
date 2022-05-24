// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main
import (
	"fmt"
	"unicode"
)

func main() {
	d := []byte("abc     a a a                   a     ccc  ddd d")
	e := RemoveSpace(d)
	fmt.Println(string(e))
}

func RemoveSpace(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			// 结合remove函数
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}